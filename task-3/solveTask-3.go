package main

import "fmt"

func solveTask3(arr []int) []int {

	last := arr[len(arr)-1]

	copy(arr[1:], arr[:len(arr)-1])
	arr[0] = last
	return arr
}

func displayArr(arr []int) {
	for _, val := range arr {
		fmt.Print(val)
		fmt.Print(" ")
	}
}
