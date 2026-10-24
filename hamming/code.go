package hamming

func Matrices(n, k int) (G, H [][]int) {
	m := n - k

	H = make([][]int, m)
	for i := range H {
		H[i] = make([]int, n)
	}

	for i := range m {
		H[i][k+i] = 1
	}

	var col int
	for value := 1; value <= n && col < k; value++ {
		if value&(value-1) == 0 {
			continue
		}

		for row := range m {
			H[row][col] = (value >> row) & 1
		}

		col++
	}

	G = make([][]int, k)
	for i := range k {
		G[i] = make([]int, n)
		G[i][i] = 1

		for j := range m {
			G[i][k+j] = H[j][i]
		}
	}

	return G, H
}
