package main

const (
	exist    = "YES"
	notExist = "NO"
)

func solveTask2(a, b, c int) string {

	if (a+b > c) && (a+c > b) && (c+b > a) {
		return exist
	}
	return notExist
}
