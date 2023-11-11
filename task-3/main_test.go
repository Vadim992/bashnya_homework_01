package main

import (
	"log"
	"reflect"
	"testing"
)

func TestSolveTask3(t *testing.T) {
	arrIn := []int{4, 5, 3, 4, 2, 3}

	expect := []int{3, 4, 5, 3, 4, 2}

	res := solveTask3(arrIn)

	if !reflect.DeepEqual(res, expect) {
		log.Fatalf("Wrong answer, expected %v", expect)
	}
}
