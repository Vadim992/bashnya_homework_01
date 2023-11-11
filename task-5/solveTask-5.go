package main

import "strings"

const (
	replaceStr = "1"
	newStr     = "one"
)

func solveTask5(str string) string {
	return strings.Replace(str, replaceStr, newStr, -1)
}
