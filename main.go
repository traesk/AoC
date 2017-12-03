package main

import (
	"fmt"
	"strconv"

	"github.com/traesk/AoC/1"
	"github.com/traesk/AoC/2"
	"github.com/traesk/AoC/3"
)

const pathDay1 = "./1/captcha"
const pathDay2 = "./2/sp"

type Solution struct {
	Day     int
	Input   string
	Answers string
}

func main() {
	day3.Solve(325489)
	day1input, day1answer, day1ext := day1.Solve(pathDay1)
	day2answer := day2.Solve1(pathDay2)
	day2answer2 := day2.Solve2(pathDay2)
	solutions := []Solution{
		{
			Day:   1,
			Input: day1input,

			Answers: day1answer + ", " + day1ext,
		},
		{
			2,
			"sp",
			strconv.Itoa(day2answer) + ", " + strconv.Itoa(day2answer2),
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
   **		stars: **

***********************************************************
***********************************************************
`)
	}

}
