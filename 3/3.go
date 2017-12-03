package day3

import (
	"io/ioutil"
)

func Solve(input int) {
	drawMemGrid(600)
}
func drawMemGrid(input int) {
	file := generate(input, 1, 0) // with digits
	ioutil.WriteFile("ulam", file, 0770)
}
