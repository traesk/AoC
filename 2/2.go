package day2

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Solve1(path string) int {
	chck, err := getChecksum(path)
	if err != nil {
		fmt.Println(":(")
		return 0
	}
	return chck
}
func Solve2(path string) int {

	chck, err := getChecksum2(path)
	if err != nil {
		fmt.Println(":(")
		return 0
	}
	return chck
}
func getChecksum2(path string) (int, error) {
	fmt.Println("getChecksum2...")
	_, spread, err := makeSpreadSheet(path)
	if err != nil {
		return 0, err
	}
	var checkSum int
	for i := 0; i < 16; i++ {
		set := spread[i]
		curr := 0
		for j := 0; j < 16; j++ {
			curr = set[j]
			// En radda nuffror står framför oss
			// Vi vill veta x / y % 0
			for h := 0; h < 16; h++ {
				temp := set[h]
				if h != j {
					if d := curr % temp; d == 0 {
						fmt.Println("Winner found!", curr, "is divisible by", temp)
						fmt.Println(curr / temp)
						checkSum += curr / temp
					}

				}
			}

		}

		checkSum += 0

	}

	return checkSum, nil
}

func getChecksum(path string) (int, error) {
	fmt.Println("getChecksum...")
	_, spread, err := makeSpreadSheet(path)
	if err != nil {
		return 0, err
	}
	var checkSum int
	for i := 0; i < 16; i++ {
		set := spread[i]
		biggest := 0
		smallest := 9999
		for j := 0; j < 16; j++ {
			// En radda nuffror står framför oss
			// Vi vill ha den största, den minsta, och skillnaden
			if set[j] < smallest {
				smallest = set[j]
			}
			if set[j] > biggest {
				biggest = set[j]
			}

		}
		fmt.Println("Set:", set)
		fmt.Println("Diff is", biggest, "-", smallest, "=", biggest-smallest)
		diff := biggest - smallest
		fmt.Println(checkSum, "+", diff, "=", checkSum+diff)
		checkSum += diff

	}

	return checkSum, nil
}
func readFile(path string) ([]byte, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func makeSpreadSheet(path string) ([]string, [16][16]int, error) {
	file, err := readFile(path)
	if err != nil {
		return nil, [16][16]int{}, err
	}

	byteStr := string(file)
	trim := strings.Trim(byteStr, "	")
	var list []string
	var seq string

	for i, b := range trim {

		s := string(b)
		_, err := strconv.Atoi(s)
		if err != nil {
			//fmt.Println("space!")
			list = append(list, seq)
			seq = ""
		} else {
			seq = seq + s
			//fmt.Println(v)
			if i == len(trim)-1 {
				list = append(list, seq)
			}
		}

	}
	//fmt.Println("len", len(list), list)
	var s [16][16]int
	counter := 0
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			s[i][j] = 0
			num, err := strconv.Atoi(list[counter])
			if counter > 616 {
				//fmt.Println("Debug")
			}
			if err != nil {
				//fmt.Println("Counter:", counter, "i:", i, "j:", j)
				counter = counter + 1
				j--
			} else {
				//fmt.Println("Add num:", num, "Counter:", counter)
				s[i][j] = num
				counter = counter + 1
			}

		}
	}

	//fmt.Println("List:", list, s)
	return list, s, nil
}
