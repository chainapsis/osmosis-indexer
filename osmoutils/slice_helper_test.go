package osmoutils_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/osmosis-labs/osmosis/osmoutils"
	"github.com/osmosis-labs/osmosis/osmoutils/osmoassert"
)

func TestReverseSlice(t *testing.T) {
	tests := map[string]struct {
		s []string

		expectedSolvedInput []string
	}{
		"Even length array":       {s: []string{"a", "b", "c", "d"}, expectedSolvedInput: []string{"d", "c", "b", "a"}},
		"Empty array":             {s: []string{}, expectedSolvedInput: []string{}},
		"Odd length array":        {s: []string{"a", "b", "c"}, expectedSolvedInput: []string{"c", "b", "a"}},
		"Single element array":    {s: []string{"a"}, expectedSolvedInput: []string{"a"}},
		"Array with empty string": {s: []string{"a", "b", "c", "", "d"}, expectedSolvedInput: []string{"d", "", "c", "b", "a"}},
		"Array with numbers":      {s: []string{"a", "b", "c", "1", "2", "3"}, expectedSolvedInput: []string{"3", "2", "1", "c", "b", "a"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualSolvedInput := osmoutils.ReverseSlice(tc.s)
			require.Equal(t, tc.expectedSolvedInput, actualSolvedInput)
		})
	}
}

func TestMergeSlices(t *testing.T) {
	lessInt := func(a, b int) bool {
		return a < b
	}
	testCases := []struct {
		name   string
		slice1 []int
		slice2 []int
		less   func(a, b int) bool
		want   []int
	}{
		{
			name:   "basic merge",
			slice1: []int{1, 3, 5},
			slice2: []int{2, 4, 6},
			less:   lessInt,
			want:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "Empty slice1",
			slice1: []int{},
			slice2: []int{2, 4, 6},
			less:   lessInt,
			want:   []int{2, 4, 6},
		},
		{
			name:   "Empty slice2",
			slice1: []int{1, 3, 5},
			slice2: []int{},
			less:   lessInt,
			want:   []int{1, 3, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := osmoutils.MergeSlices(tc.slice1, tc.slice2, lessInt)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestContainsDuplicateDeepEqual(t *testing.T) {
	tests := []struct {
		input []interface{}
		want  bool
	}{
		{[]interface{}{[]int{1, 2, 3}, []int{4, 5, 6}}, false},
		{[]interface{}{[]int{1, 2, 3}, []int{1, 2, 3}}, true},
		{[]interface{}{[]string{"hello", "world"}, []string{"goodbye", "world"}}, false},
		{[]interface{}{[]string{"hello", "world"}, []string{"hello", "world"}}, true},
		{[]interface{}{[][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}}, true},
	}

	for _, tt := range tests {
		got := osmoutils.ContainsDuplicateDeepEqual(tt.input)
		require.Equal(t, tt.want, got)
	}
}

func TestIThSmallest(t *testing.T) {
	var (
		less = func(a, b int) bool {
			return a < b
		}
		greater = func(a, b int) bool {
			return a > b
		}
	)

	tests := map[string]struct {
		s           []int
		i           int
		less        osmoutils.LessFunc[int]
		expected    int
		expectPanic bool
	}{
		"one element": {
			s:        []int{1},
			i:        0,
			less:     less,
			expected: 1,
		},
		"five elements, sorted": {
			s:        []int{1, 2, 3, 4, 5},
			i:        3,
			less:     less,
			expected: 4,
		},
		"five elements, unsorted, smallest i": {
			s:        []int{4, 21, 3, 1, 10},
			i:        0,
			less:     less,
			expected: 1,
		},
		"five elements, unsorted, largest i": {
			s:        []int{4, 21, 3, 1, 10},
			i:        4,
			less:     less,
			expected: 21,
		},
		"five elements, repeated elements": {
			s:        []int{4, 21, 3, 3, 10},
			i:        2,
			less:     less,
			expected: 4,
		},
		"many elements, unsorted, odd number of medians, smallest": {
			// this sorted is:
			// 1, 3, 4, 10, 12, 13, 13, 14, 21, 23, 45, 48, 86, 98, 100
			s: []int{
				4, 21, 3, 1, 10,
				98, 23, 13, 48, 12,
				100, 45, 14, 86, 13,
			},
			i:        10,
			less:     less,
			expected: 45,
		},
		"many elements, unsorted, even number of medians, smallest": {
			// this sorted is:
			// 1, 3, 4, 10, 12, 13, 21, 23, 48, 98
			s: []int{
				4, 21, 3, 1, 10,
				98, 23, 13, 48, 12,
			},
			i:        3,
			less:     less,
			expected: 10,
		},
		"number of elements > 5, requires search for pivot": {
			// this sorted is:
			// 1, 1, 3, 3, 4, 4, 10, 10, 12, 12, 13, 13, 13, 13, 14, 14, 21, 21, 23, 23, 45, 45, 48, 48, 86, 86, 98, 98, 100, 100
			s: []int{
				4, 21, 3, 1, 10,
				98, 23, 13, 48, 12,
				100, 45, 14, 86, 13,
				4, 21, 3, 1, 10,
				98, 23, 13, 48, 12,
				100, 45, 14, 86, 13,
			},
			i:        20,
			less:     less,
			expected: 45,
		},
		"many elements, greater function": {
			// this sorted is:
			// 1, 3, 4, 10, 12, 13, 13, 14, 21, 23, 45, 48, 86, 98, 100
			s: []int{
				4, 21, 3, 1, 10,
				98, 23, 13, 48, 12,
				100, 45, 14, 86, 13,
			},
			i:        10,
			less:     greater,
			expected: 12,
		},
		"panic: empty": {
			s:           []int{},
			i:           0,
			less:        less,
			expected:    0,
			expectPanic: true,
		},
		"panic: out of bounds": {
			s:           []int{1, 2},
			i:           3,
			less:        less,
			expected:    0,
			expectPanic: true,
		},
		"panic: negative i": {
			s:           []int{1, 2},
			i:           -1,
			less:        less,
			expected:    0,
			expectPanic: true,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			osmoassert.ConditionalPanic(t, tc.expectPanic, func() {
				result := osmoutils.IThSmallest(tc.s, tc.i, tc.less)
				require.Equal(t, tc.expected, result)
			})
		})
	}
}
