package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b, c int

	fmt.Fscan(in, &a, &b, &c)

	fmt.Println(solveTask2(a, b, c))
}
