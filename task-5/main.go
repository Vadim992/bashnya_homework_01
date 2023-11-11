package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var str string

	fmt.Fscan(in, &str)

	fmt.Println(solveTask5(str))

}
