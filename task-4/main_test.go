package main

import (
	"log"
	"testing"
)

func TestSolveTask4(t *testing.T) {
	n := 3
	matrIn := [][]int{
		{0, 1, 2},
		{1, 5, 3},
		{2, 3, 4},
	}

	expect := sim

	res := solveTask4(n, matrIn)

	if res != expect {
		log.Fatalf("Wrang answer, expected %s", expect)
	}

	matrIn = [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{1, 0, 0},
	}

	expect = noSim

	res = solveTask4(n, matrIn)
	if res != expect {
		log.Fatalf("Wrang answer, expected %s", expect)
	}
}
