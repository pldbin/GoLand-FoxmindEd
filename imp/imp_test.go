package implementminimumandmaximumvalues

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxMinSum(t *testing.T) {
	req := require.New(t)

	test := func(min, max, minsum, maxsum int) func(t *testing.T) {
		return func(t *testing.T) {
			min, max, minsum, maxsum := MaxMinSum([5]int{})
			req.Equal(min, max, minsum, maxsum)
		}
	}
	t.Run("ok", test(2, 21, 41, 60))
	t.Run("no", test(0, 0, 0, 0))
}

//func TestMaxMinSum(t *testing.T) {
//	type args struct {
//		arr [5]int
//	}
//	tests := []struct {
//		name       string
//		args       args
//		wantMin    int
//		wantMax    int
//		wantMinsum int
//		wantMaxsum int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			gotMin, gotMax, gotMinsum, gotMaxsum := MaxMinSum(tt.args.arr)
//			if gotMin != tt.wantMin {
//				t.Errorf("MaxMinSum() gotMin = %v, want %v", gotMin, tt.wantMin)
//			}
//			if gotMax != tt.wantMax {
//				t.Errorf("MaxMinSum() gotMax = %v, want %v", gotMax, tt.wantMax)
//			}
//			if gotMinsum != tt.wantMinsum {
//				t.Errorf("MaxMinSum() gotMinsum = %v, want %v", gotMinsum, tt.wantMinsum)
//			}
//			if gotMaxsum != tt.wantMaxsum {
//				t.Errorf("MaxMinSum() gotMaxsum = %v, want %v", gotMaxsum, tt.wantMaxsum)
//			}
//		})
//	}
//}
