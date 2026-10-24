package matrix

func Transpose(A [][]int) [][]int {
	rows, cols := len(A), len(A[0])

	T := make([][]int, cols)
	for j := range cols {
		T[j] = make([]int, rows)
		for i := range rows {
			T[j][i] = A[i][j]
		}
	}

	return T
}

func MatMul(A, B [][]int) [][]int {
	rowsA, colsA := len(A), len(A[0])
	rowsB, colsB := len(B), len(B[0])

	if colsA != rowsB {
		panic("matrix dimensions do not match")
	}

	C := make([][]int, rowsA)
	for i := range rowsA {
		C[i] = make([]int, colsB)
		for j := range colsB {
			for k := range colsA {
				C[i][j] ^= A[i][k] & B[k][j]
			}
		}
	}

	return C
}

func Apply(H [][]int, v []int) []int {
	C := make([]int, len(H))
	for i := range H {
		for j := range v {
			C[i] ^= H[i][j] & v[j]
		}
	}

	return C
}

func IsZero(A [][]int) bool {
	for i := range A {
		for j := range A[i] {
			if A[i][j] != 0 {
				return false
			}
		}
	}

	return true
}
