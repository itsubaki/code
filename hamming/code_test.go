package hamming_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/itsubaki/code/hamming"
	"github.com/itsubaki/code/matrix"
)

func ExampleMatrices() {
	G, H := hamming.Matrices(7, 4)

	for _, row := range G {
		fmt.Println(row)
	}
	fmt.Println(".")

	for _, row := range H {
		fmt.Println(row)
	}
	fmt.Println(".")

	GHT := matrix.MatMul(G, matrix.Transpose(H))
	fmt.Println(matrix.IsZero(GHT))

	// Output:
	// [1 0 0 0 1 1 0]
	// [0 1 0 0 1 0 1]
	// [0 0 1 0 0 1 1]
	// [0 0 0 1 1 1 1]
	// .
	// [1 1 0 1 1 0 0]
	// [1 0 1 1 0 1 0]
	// [0 1 1 1 0 0 1]
	// .
	// true
}

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
		r    []int
		want []int // syndrome
	}{
		// 0L
		{
			r:    []int{0, 0, 0, 0, 0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{1, 0, 1, 0, 1, 0, 1},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{0, 1, 1, 0, 0, 1, 1},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{1, 1, 0, 0, 1, 1, 0},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{0, 0, 0, 1, 1, 1, 1},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{1, 0, 1, 1, 0, 1, 0},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{0, 1, 1, 1, 1, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{1, 1, 0, 1, 0, 0, 1},
			want: []int{0, 0, 0},
		},
		// 1L
		{
			r:    []int{1, 1, 1, 1, 1, 1, 1},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{0, 1, 0, 1, 0, 1, 0},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{1, 0, 0, 1, 1, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{0, 0, 1, 1, 0, 0, 1},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{1, 1, 1, 0, 0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{0, 1, 0, 0, 1, 0, 1},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{1, 0, 0, 0, 0, 1, 1},
			want: []int{0, 0, 0},
		},
		{
			r:    []int{0, 0, 1, 0, 1, 1, 0},
			want: []int{0, 0, 0},
		},
	}

	for _, c := range cases {
		syndrome := matrix.Apply(H, c.r)
		if !slices.Equal(syndrome, c.want) {
			t.Errorf("syndrome mismatch: got %v, want %v",
				syndrome, c.want)
		}
	}
}
