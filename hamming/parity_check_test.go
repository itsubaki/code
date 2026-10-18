package hamming_test

import (
	"slices"
	"testing"
)

func Test_parityCheckMatrix(t *testing.T) {
	H := [][]int{
		{0, 0, 0, 1, 1, 1, 1},
		{0, 1, 1, 0, 0, 1, 1},
		{1, 0, 1, 0, 1, 0, 1},
	}

	cases := []struct {
		r        []int
		syndrome []int
	}{
		{
			r:        []int{0, 0, 0, 0, 0, 0, 0},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{1, 0, 0, 0, 0, 0, 0},
			syndrome: []int{0, 0, 1},
		},
		{
			r:        []int{0, 1, 0, 0, 0, 0, 0},
			syndrome: []int{0, 1, 0},
		},
		{
			r:        []int{0, 0, 1, 0, 0, 0, 0},
			syndrome: []int{0, 1, 1},
		},
		{
			r:        []int{0, 0, 0, 1, 0, 0, 0},
			syndrome: []int{1, 0, 0},
		},
		{
			r:        []int{0, 0, 0, 0, 1, 0, 0},
			syndrome: []int{1, 0, 1},
		},
		{
			r:        []int{0, 0, 0, 0, 0, 1, 0},
			syndrome: []int{1, 1, 0},
		},
		{
			r:        []int{0, 0, 0, 0, 0, 0, 1},
			syndrome: []int{1, 1, 1},
		},
		{
			r:        []int{1, 0, 1, 1, 0, 1, 0},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{0, 0, 0, 1, 1, 1, 1},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{1, 0, 1, 1, 0, 1, 1},
			syndrome: []int{1, 1, 1},
		},
	}

	for _, c := range cases {
		syndrome := make([]int, len(H))
		for i := range H {
			for j := range c.r {
				syndrome[i] ^= H[i][j] & c.r[j]
			}
		}

		if !slices.Equal(syndrome, c.syndrome) {
			t.Errorf("syndrome mismatch: got %v, want %v",
				syndrome, c.syndrome)
		}
	}
}
