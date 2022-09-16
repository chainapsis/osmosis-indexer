package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/osmosis-labs/osmosis/v12/x/gamm/types"
)

// EstimateMultihopSwapExactAmountIn iterates `EstimateSwapExactAmountIn` and returns
// the total token out given route for multihop.
func (k Keeper) EstimateMultihopSwapExactAmountIn(
	ctx sdk.Context,
	routes []types.SwapAmountInRoute,
	tokenIn sdk.Coin,
	tokenOutMinAmount sdk.Int,
) (tokenOutAmount sdk.Int, err error) {
	for i, route := range routes {
		swapFeeMultiplier := sdk.OneDec()
		if types.SwapAmountInRoutes(routes).IsOsmoRoutedMultihop() {
			swapFeeMultiplier = types.MultihopSwapFeeMultiplierForOsmoPools.Clone()
		}

		// To prevent the multihop swap from being interrupted prematurely, we keep
		// the minimum expected output at a very low number until the last pool
		_outMinAmount := sdk.NewInt(1)
		if len(routes)-1 == i {
			_outMinAmount = tokenOutMinAmount
		}

		// Execute the expected swap on the current routed pool
		pool, poolErr := k.getPoolForSwap(ctx, route.PoolId)
		if poolErr != nil {
			return sdk.Int{}, poolErr
		}

		swapFee := pool.GetSwapFee(ctx).Mul(swapFeeMultiplier)

		// Execute the expected swap on the current routed pool
		tokenOutAmount, err = k.EstimateSwapExactAmountIn(ctx, route.PoolId, tokenIn, route.TokenOutDenom, _outMinAmount, swapFee)
		if err != nil {
			return sdk.Int{}, err
		}

		// Chain output of current pool as the input for the next routed pool
		tokenIn = sdk.NewCoin(route.TokenOutDenom, tokenOutAmount)
	}
	return tokenOutAmount, nil
}

// EstimateSwapExactAmountIn estimates the amount of token out given the exact amount of token in.
// This method does not execute the full steps of an actaul swap,
// but estimates the amount of token out by only manipulating the state of the pool.
// This method should only be called by query methods.
func (k Keeper) EstimateSwapExactAmountIn(
	ctx sdk.Context,
	poolId uint64,
	tokenIn sdk.Coin,
	tokenOutDenom string,
	tokenOutMinAmount sdk.Int,
	swapFee sdk.Dec,
) (sdk.Int, error) {
	pool, err := k.getPoolForSwap(ctx, poolId)
	if err != nil {
		return sdk.Int{}, err
	}

	_, tokenOut, err := k.estimateSwapExactAmountIn(ctx, pool, tokenIn, tokenOutDenom, tokenOutMinAmount, swapFee)
	if err != nil {
		return tokenOut.Amount, err
	}

	err = k.estimateUpdatePoolForSwap(ctx, pool, tokenIn, tokenOut)
	return tokenOut.Amount, err
}

// estimateSwapExactAmountIn performs `SwapOutAmtGivenIn` for the given inputs.
// Note that this method does not include writing new state to the store, but only has the
// logic for calculation for the swap.
func (k Keeper) estimateSwapExactAmountIn(
	ctx sdk.Context,
	pool types.PoolI,
	tokenIn sdk.Coin,
	tokenOutDenom string,
	tokenOutMinAmount sdk.Int,
	swapFee sdk.Dec,
) (updatedPool types.PoolI, tokenOut sdk.Coin, err error) {
	if tokenIn.Denom == tokenOutDenom {
		return pool, sdk.Coin{}, errors.New("cannot trade same denomination in and out")
	}
	tokensIn := sdk.Coins{tokenIn}

	// Executes the swap in the pool and stores the output. Updates pool assets but
	// does not actually transfer any tokens to or from the pool.
	tokenOutCoin, err := pool.SwapOutAmtGivenIn(ctx, tokensIn, tokenOutDenom, swapFee)
	if err != nil {
		return pool, sdk.Coin{}, err
	}

	tokenOutAmount := tokenOutCoin.Amount

	if !tokenOutAmount.IsPositive() {
		return pool, sdk.Coin{}, sdkerrors.Wrapf(types.ErrInvalidMathApprox, "token amount must be positive")
	}

	if tokenOutAmount.LT(tokenOutMinAmount) {
		return pool, sdk.Coin{}, sdkerrors.Wrapf(types.ErrLimitMinAmount, "%s token is lesser than min amount", tokenOutDenom)
	}

	return pool, tokenOutCoin, nil
}

// EstimateMultihopSwapExactAmountOut iterates `EstimateSwapExactAmountOut` and returns
// the total token in given route for multihop.
func (k Keeper) EstimateMultihopSwapExactAmountOut(
	ctx sdk.Context,
	routes []types.SwapAmountOutRoute,
	tokenInMaxAmount sdk.Int,
	tokenOut sdk.Coin,
) (tokenInAmount sdk.Int, err error) {
	swapFeeMultiplier := sdk.OneDec()

	if types.SwapAmountOutRoutes(routes).IsOsmoRoutedMultihop() {
		swapFeeMultiplier = types.MultihopSwapFeeMultiplierForOsmoPools.Clone()
	}

	// Determine what the estimated input would be for each pool along the multihop route
	insExpected, err := k.createMultihopExpectedSwapOuts(ctx, routes, tokenOut, swapFeeMultiplier)
	if err != nil {
		return sdk.Int{}, err
	}
	if len(insExpected) == 0 {
		return sdk.Int{}, nil
	}

	insExpected[0] = tokenInMaxAmount

	// Iterates through each routed pool and executes their respective swaps. Note that all of the work to get the return
	// value of this method is done when we calculate insExpected – this for loop primarily serves to execute the actual
	// swaps on each pool.
	for i, route := range routes {
		_tokenOut := tokenOut

		// If there is one pool left in the route, set the expected output of the current swap
		// to the estimated input of the final pool.
		if i != len(routes)-1 {
			_tokenOut = sdk.NewCoin(routes[i+1].TokenInDenom, insExpected[i+1])
		}

		// Execute the expected swap on the current routed pool
		pool, poolErr := k.getPoolForSwap(ctx, route.PoolId)
		if poolErr != nil {
			return sdk.Int{}, poolErr
		}
		swapFee := pool.GetSwapFee(ctx).Mul(swapFeeMultiplier)
		// Execute the expected swap on the current routed pool
		_tokenInAmount, err := k.EstimateSwapExactAmountOut(ctx, route.PoolId, route.TokenInDenom, insExpected[i], _tokenOut, swapFee)
		if err != nil {
			return sdk.Int{}, err
		}

		// Sets the final amount of tokens that need to be input into the first pool. Even though this is the final return value for the
		// whole method and will not change after the first iteration, we still iterate through the rest of the pools to execute their respective
		// swaps.
		if i == 0 {
			tokenInAmount = _tokenInAmount
		}
	}

	return tokenInAmount, nil
}

// EstimateSwapExactAmountOut estimates the amount of token out given the exact amount of token in.
// This method does not execute the full steps of an actaul swap,
// but estimates the amount of token out by only manipulating the state of the pool.
// This method should only be called by query methods.
func (k Keeper) EstimateSwapExactAmountOut(
	ctx sdk.Context,
	poolId uint64,
	tokenInDenom string,
	tokenInMaxAmount sdk.Int,
	tokenOut sdk.Coin,
	swapFee sdk.Dec,
) (tokenInAmount sdk.Int, err error) {
	pool, err := k.getPoolForSwap(ctx, poolId)
	if err != nil {
		return sdk.Int{}, err
	}

	_, tokenIn, err := k.estimateSwapExactAmountOut(ctx, pool, tokenInDenom, tokenInMaxAmount, tokenOut, swapFee)
	if err != nil {
		return tokenIn.Amount, err
	}

	err = k.estimateUpdatePoolForSwap(ctx, pool, tokenIn, tokenOut)
	return tokenIn.Amount, err
}

func (k Keeper) estimateSwapExactAmountOut(
	ctx sdk.Context,
	pool types.PoolI,
	tokenInDenom string,
	tokenInMaxAmount sdk.Int,
	tokenOut sdk.Coin,
	swapFee sdk.Dec,
) (updatedPool types.PoolI, tokenIn sdk.Coin, err error) {
	if tokenInDenom == tokenOut.Denom {
		return pool, sdk.Coin{}, errors.New("cannot trade same denomination in and out")
	}

	poolOutBal := pool.GetTotalPoolLiquidity(ctx).AmountOf(tokenOut.Denom)
	if tokenOut.Amount.GTE(poolOutBal) {
		return pool, sdk.Coin{}, sdkerrors.Wrapf(types.ErrTooManyTokensOut,
			"can't get more tokens out than there are tokens in the pool")
	}

	tokenIn, err = pool.SwapInAmtGivenOut(ctx, sdk.Coins{tokenOut}, tokenInDenom, swapFee)
	if err != nil {
		return pool, sdk.Coin{}, err
	}
	tokenInAmount := tokenIn.Amount

	if tokenInAmount.LTE(sdk.ZeroInt()) {
		return pool, sdk.Coin{}, sdkerrors.Wrapf(types.ErrInvalidMathApprox, "token amount is zero or negative")
	}

	if tokenInAmount.GT(tokenInMaxAmount) {
		return pool, sdk.Coin{}, sdkerrors.Wrapf(types.ErrLimitMaxAmount, "Swap requires %s, which is greater than the amount %s", tokenIn, tokenInMaxAmount)
	}

	return pool, tokenIn, nil
}

// estimateUpdatePoolForSwap updates the pool balance.
// This method does not include emitting hooks or sending the actual asset from
// the sender to the pool or vice versa, as this is designed to be used by query methods.
func (k Keeper) estimateUpdatePoolForSwap(
	ctx sdk.Context,
	pool types.PoolI,
	tokenIn sdk.Coin,
	tokenOut sdk.Coin,
) error {
	tokensIn := sdk.Coins{tokenIn}
	tokensOut := sdk.Coins{tokenOut}

	err := k.setPool(ctx, pool)
	if err != nil {
		return err
	}

	k.RecordTotalLiquidityIncrease(ctx, tokensIn)
	k.RecordTotalLiquidityDecrease(ctx, tokensOut)
	return err
}
