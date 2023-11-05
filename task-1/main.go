package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var a, b float64

	fmt.Fscan(in, &a, &b)
	fmt.Println(solveTask1(a, b))
}
