package main

import (
	"log"
	"testing"
)

func TestSolverTask1(t *testing.T) {
	var a, b float64 = 3, 4

	expect := float64(5)
	res := solveTask1(a, b)

	if res != expect {
		log.Fatalf("Wrong answer, expected %v", expect)
	}
}
