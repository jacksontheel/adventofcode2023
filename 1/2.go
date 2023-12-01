package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r, _ := regexp.Compile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		start := 0
		digits := make([]string, 0, len(line))
		for start < len(line) {
			digitIndex := r.FindStringSubmatchIndex(line[start:])

			if len(digitIndex) > 0 {
				digits = append(digits, line[start+digitIndex[0]:start+digitIndex[1]])

				// If the matched string is not a digit set digitIndex to be one further back
				// to catch overlap
				if _, e := strconv.Atoi(line[start+digitIndex[0] : start+digitIndex[1]]); e == nil {
					start += digitIndex[1]
				} else {
					start += digitIndex[1] - 1
				}
			} else {
				start = len(line)
			}
		}

		if len(digits) == 1 {
			fmt.Println(digits[0])
			add := 0
			if n := getNumberString(digits[0]); n != 0 {
				add = n
			} else {
				add, _ = strconv.Atoi(digits[0])
			}
			sum += add * 11
		} else {
			fmt.Println(digits[0] + " " + digits[len(digits)-1])
			add1 := 0
			add2 := 0

			if n := getNumberString(digits[0]); n != 0 {
				add1 = n
			} else {
				add1, _ = strconv.Atoi(digits[0])
			}

			if n := getNumberString(digits[len(digits)-1]); n != 0 {
				add2 = n
			} else {
				add2, _ = strconv.Atoi(digits[len(digits)-1])
			}

			sum += add1*10 + add2
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(sum)
	}
}

func getNumberString(s string) int {
	switch s {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	return 0
}
