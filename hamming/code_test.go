package hamming_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/itsubaki/code/matrix"
)

func Example() {
	H := [][]int{
		{0, 0, 0, 1, 1, 1, 1},
		{0, 1, 1, 0, 0, 1, 1},
		{1, 0, 1, 0, 1, 0, 1},
	}

	G := [][]int{
		{1, 0, 0, 0, 0, 1, 1},
		{0, 1, 0, 0, 1, 0, 1},
		{0, 0, 1, 0, 1, 1, 0},
		{0, 0, 0, 1, 1, 1, 1},
	}

	GHT := matrix.MatMul(G, matrix.Transpose(H))
	fmt.Println(matrix.IsZero(GHT))

	// Output:
	// true
}

func Test_steane(t *testing.T) {
	H := [][]int{
		{0, 0, 0, 1, 1, 1, 1},
		{0, 1, 1, 0, 0, 1, 1},
		{1, 0, 1, 0, 1, 0, 1},
	}

	cases := []struct {
		r        []int
		syndrome []int
	}{
		// 0L
		{
			r:        []int{0, 0, 0, 0, 0, 0, 0},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{1, 0, 1, 0, 1, 0, 1},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{0, 1, 1, 0, 0, 1, 1},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{1, 1, 0, 0, 1, 1, 0},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{0, 0, 0, 1, 1, 1, 1},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{1, 0, 1, 1, 0, 1, 0},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{0, 1, 1, 1, 1, 0, 0},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{1, 1, 0, 1, 0, 0, 1},
			syndrome: []int{0, 0, 0},
		},
		// 1L
		{
			r:        []int{1, 1, 1, 1, 1, 1, 1},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{0, 1, 0, 1, 0, 1, 0},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{1, 0, 0, 1, 1, 0, 0},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{0, 0, 1, 1, 0, 0, 1},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{1, 1, 1, 0, 0, 0, 0},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{0, 1, 0, 0, 1, 0, 1},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{1, 0, 0, 0, 0, 1, 1},
			syndrome: []int{0, 0, 0},
		},
		{
			r:        []int{0, 0, 1, 0, 1, 1, 0},
			syndrome: []int{0, 0, 0},
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
