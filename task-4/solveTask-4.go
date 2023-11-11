package main

const (
	sim   = "yes"
	noSim = "no"
)

func solveTask4(n int, matr [][]int) string {

	for i := 0; i < n; i++ {

		for j := i + 1; j < n; j++ {
			if matr[i][j] != matr[j][i] {
				return noSim
			}
		}
	}

	return sim
}
