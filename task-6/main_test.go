package main

import (
	"log"
	"testing"
)

func TestSolveTask6(t *testing.T) {
	in := "1 2 3 2 1"

	expect := 3

	res := solveTask6(in)

	if res != expect {
		log.Fatalf("Wrong answer, expected %d", expect)
	}

	in = "1 2 3 4 5 6 7 8 9 10"

	expect = 10

	res = solveTask6(in)

	if res != expect {
		log.Fatalf("Wrong answer, expected %d", expect)
	}

	in = "1 2 3 4 5 1 2 1 2 7 3"

	expect = 6

	res = solveTask6(in)

	if res != expect {
		log.Fatalf("Wrong answer, expected %d", expect)
	}
}
