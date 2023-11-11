package main

import (
	"log"
	"strings"
	"testing"
)

func TestSolveTask5(t *testing.T) {

	in := "1+1=2"

	expect := "one+one=2"

	res := solveTask5(in)

	if strings.Compare(res, expect) != 0 {
		log.Fatalf("Wrong answer, expected %s", expect)
	}
}
