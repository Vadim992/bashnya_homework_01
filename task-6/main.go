package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	line = strings.TrimSpace(line)

	fmt.Println(solveTask6(line))
}
