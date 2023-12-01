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

	r, _ := regexp.Compile(`\d`)
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		digits := r.FindAllString(scanner.Text(), -1)

		if len(digits) == 1 {
			add, _ := strconv.Atoi(digits[0])
			sum += add * 11
		} else {
			add1, _ := strconv.Atoi(digits[0])
			add2, _ := strconv.Atoi(digits[len(digits)-1])
			sum += add1*10 + add2
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(sum)
	}
}
