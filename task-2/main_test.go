package main

import (
	"log"
	"testing"
)

func TestSolveTask2(t *testing.T) {
	a, b, c := 3, 4, 5
	expect := exist
	res := solveTask2(a, b, c)

	if res != expect {
		log.Fatalf("Wrong answer, expected %s", expect)
	}
}
