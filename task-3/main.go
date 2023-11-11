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

	s := make([]int, n)

	for i := 0; i < n; i++ {
		var num int

		fmt.Fscan(in, &num)

		s[i] = num
	}

	displayArr(solveTask3(s))

}
