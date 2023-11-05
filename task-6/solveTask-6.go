package main

import "strings"

func solveTask6(line string) int {
	s := strings.Split(line, " ")

	numMap := make(map[string]struct{}, 100000)

	for _, val := range s {

		if _, ok := numMap[val]; !ok {
			numMap[val] = struct{}{}
		}
	}

	return len(numMap)
}
