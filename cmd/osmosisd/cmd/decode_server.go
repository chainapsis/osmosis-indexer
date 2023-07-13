package cmd

// DONTCOVER

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/osmosis-labs/osmosis/v16/app"
	"github.com/osmosis-labs/osmosis/v16/cmd/osmosisd/cmd/decoder"
)

const (
	decodeServerPort = "decodeServerPort"
)

// decoderCmd gets cmd to run decode server
func decoderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "decoder",
		Short: "Example osmosisd decode -p 8888, which would run decoder server on specified port",
		Long:  "decoder command runs decoder server to decode byte array to General Cosmos messages",
		RunE: func(cmd *cobra.Command, args []string) error {
			decodeServerFlag, err := cmd.Flags().GetString(decodeServerPort)
			if err != nil {
				return err
			}

			d := decoder.Decoder{
				EncodingConfig: app.MakeEncodingConfig(),
			}
			err = d.ListenAndServe(decodeServerFlag)
			if err != nil {
				fmt.Println(err)
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringP(decodeServerPort, "p", "", "port to listen to")
	cmd.MarkFlagRequired(decodeServerPort)
	return cmd
}
