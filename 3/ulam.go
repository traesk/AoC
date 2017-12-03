package day3

import (
	"fmt"
	"math"
)

type Direction byte

const (
	RIGHT Direction = iota
	UP
	LEFT
	DOWN
)

func generate(n, i int, c byte) []byte {
	var blob []byte
	s := make([][]string, n)
	for i := 0; i < n; i++ {
		s[i] = make([]string, n)
	}
	dir := RIGHT
	y := n / 2
	var x int
	if n%2 == 0 {
		x = y - 1
	} else {
		x = y
	} // shift left for even n's

	for j := i; j <= n*n-1+i; j++ {

		if isPrime(j) {
			str := fmt.Sprintf("%3d", c)
			str = leftPad(str)
			cb := []byte(str)
			if c == 0 {
				str := fmt.Sprintf("%3d", j)
				str = leftPad(str)
				s[y][x] = fmt.Sprintf(str)
				//blob = append(blob, []byte(fmt.Sprintf(" %3d kiss", j))...)
			} else {
				str := fmt.Sprintf("%3d", c)
				str = leftPad(str)
				s[y][x] = fmt.Sprintf(str)
				blob = append(blob, cb...)
			}
		} else {
			str := fmt.Sprintf("%3d", j)
			str = leftPad(str)
			s[y][x] = str
			//blob = append(blob, b...)
		}

		switch dir {
		case RIGHT:
			if x <= n-1 && s[y-1][x] == "" && j > i {
				dir = UP
			}
		case UP:
			if s[y][x-1] == "" {
				dir = LEFT
			}
		case LEFT:
			if x == 0 || s[y+1][x] == "" {
				dir = DOWN
			}
		case DOWN:
			if s[y][x+1] == "" {
				dir = RIGHT
			}
		}

		switch dir {
		case RIGHT:
			x += 1
		case UP:
			y -= 1
		case LEFT:
			x -= 1
		case DOWN:
			y += 1
		}
		if j == 0 {

		}
	}

	for _, row := range s {
		newRow := fmt.Sprintf("%v \n", row)
		fmt.Println(newRow)
		blob = append(blob, []byte(newRow)...)
	}
	fmt.Println()
	return blob
}

func isPrime(a int) bool {
	if a == 2 {
		return true
	}
	if a <= 1 || a%2 == 0 {
		return false
	}
	max := int(math.Sqrt(float64(a)))
	for n := 3; n <= max; n += 2 {
		if a%n == 0 {
			return false
		}
	}
	return true
}

func main() {
	generate(9, 1, 0)   // with digits
	generate(9, 1, '*') // with *
}

func leftPad(s string) string {
	if len(s) < 2 {
		return s + "      "
	} else if len(s) < 3 {
		return s + "    "
	} else if len(s) < 4 {
		return s + "   "
	} else if len(s) < 5 {
		return s + "  "
	} else if len(s) < 6 {
		return s + " "
	}
	return s
}
