package main

import (
	"AoC/1"
	"fmt"
	"strconv"
)

const pathDay1 = "./1/captcha"

type Solution struct {
	Day     int
	Input   string
	Answers string
}

func main() {
	day1input, day1answer, day1ext := day1.Solve(pathDay1)
	solutions := []Solution{
		{
			Day:   1,
			Input: day1input,

			Answers: day1answer + ", " + day1ext,
		},
	}

	for _, s := range solutions {
		fmt.Print(`
***********************************************************
***********************************************************

    *
   ***		day: ` + strconv.Itoa(s.Day) + `
  *****		input: ` + s.Input + `
 *******
*********	answer: ` + s.Answers + `
   **

***********************************************************
***********************************************************
`)
	}

}
