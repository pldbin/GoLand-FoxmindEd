package __2_imp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxMinSum(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		test1 := [5]int{5, 2, 14, 21, 20}

		expectMin := 2
		expectMax := 21
		expectMinSum := 62
		expectMaxSum := 62

		min, max, minSum, maxSum := MaxMinSum(test1)
		require.Equal(t, expectMin, min)
		require.Equal(t, expectMax, max)
		require.Equal(t, expectMinSum, minSum)
		require.Equal(t, expectMaxSum, maxSum)
	})
	t.Run("tets2", func(t *testing.T) {
		test2 := [5]int{1, 2, 3, 4, 5}

		expectMin := 1
		expectMax := 5
		expectMinSum := 15
		expectMaxSum := 15

		min, max, minSum, maxSum := MaxMinSum(test2)
		require.Equal(t, expectMin, min)
		require.Equal(t, expectMax, max)
		require.Equal(t, expectMinSum, minSum)
		require.Equal(t, expectMaxSum, maxSum)
	})
	t.Run("tets3", func(t *testing.T) {
		test3 := [5]int{21, 32, 43, 54, 11}

		expectMin := 0
		expectMax := 0
		expectMinSum := 0
		expectMaxSum := 0

		min, max, minSum, maxSum := MaxMinSum(test3)
		require.Equal(t, expectMin, min)
		require.Equal(t, expectMax, max)
		require.Equal(t, expectMinSum, minSum)
		require.Equal(t, expectMaxSum, maxSum)
	})

}

func TestMaxMinSum1(t *testing.T) {
	type args struct {
		arr [5]int
	}
	tests := []struct {
		name       string
		args       args
		wantMin    int
		wantMax    int
		wantMinsum int
		wantMaxsum int
	}{
		{
			name:       "expect from arr [5]int{5,2,14,21,20 correct conclusion Min:2,Max:21,Minsum:62,Maxsum:62}",
			args:       args{arr: [5]int{5, 2, 14, 21, 20}},
			wantMin:    2,
			wantMax:    21,
			wantMinsum: 62,
			wantMaxsum: 62,
		},
		{
			name:       "expect from arr [5]int{5,2,14,21,20 broken test Min:2-want 21,Max:21-want 22,Minsum:62-want 61,Maxsum:62-want 63",
			args:       args{arr: [5]int{5, 2, 14, 21, 20}},
			wantMin:    21,
			wantMax:    22,
			wantMinsum: 61,
			wantMaxsum: 63,
		},
		{
			name:       "expect from arr [5]int{5,2,14,21,20 correct conclusion Min:3,Max:24,Minsum:70,Maxsum:70",
			args:       args{arr: [5]int{6, 3, 15, 22, 24}},
			wantMin:    3,
			wantMax:    24,
			wantMinsum: 70,
			wantMaxsum: 70,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax, gotMinsum, gotMaxsum := MaxMinSum(tt.args.arr)
			require.Equal(t, tt.wantMin, gotMin)
			require.Equal(t, tt.wantMax, gotMax)
			require.Equal(t, tt.wantMinsum, gotMinsum)
			require.Equal(t, tt.wantMaxsum, gotMaxsum)

		})
	}
}
