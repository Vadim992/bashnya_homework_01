package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int

	fmt.Fscan(in, &n)

	matr := make([][]int, 0, n)

	for i := 0; i < n; i++ {
		s := make([]int, 0, n)

		for j := 0; j < n; j++ {

			var num int

			fmt.Fscan(in, &num)

			s = append(s, num)
		}

		matr = append(matr, s)
	}

	fmt.Println(solveTask4(n, matr))
}
