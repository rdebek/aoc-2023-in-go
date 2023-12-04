package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func replaceStringNumbersWithActualNumbers(line string) string {
	digitMapping := map[string]string{"one": "one1one", "two": "two2two", "three": "three3three", "four": "four4four", "five": "five5five", "six": "six6six", "seven": "seven7seven", "eight": "eight8eight", "nine": "nine9nine"}
	for k, v := range digitMapping {
		line = strings.Replace(line, k, v, -1)
	}
	return line
}

func getCombinedNumber(numbers []string) int {
	combinedNumber, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
	return combinedNumber
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	splitData := strings.Split(string(dat), "\n")
	var result int
	for _, line := range splitData {
		line = replaceStringNumbersWithActualNumbers(line)
		var numbers []string
		for _, character := range line {
			if _, err := strconv.Atoi(string(character)); err == nil {
				numbers = append(numbers, string(character))
			}
		}
		NumbersSumForRow := getCombinedNumber(numbers)
		fmt.Println(NumbersSumForRow)
		result += NumbersSumForRow
	}
	fmt.Print(result)
}
