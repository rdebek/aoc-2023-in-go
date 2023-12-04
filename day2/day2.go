package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func extractColor(line string, color string) []string {
	r, _ := regexp.Compile(`(\d+)\s+` + color)
	matches := r.FindAllStringSubmatch(line, -1)
	var digits []string
	for _, match := range matches {
		digits = append(digits, match[1])
	}
	return digits
}

func isGameValid(line string) bool{
	maxNumberOfColors := map[string]int{"red": 12, "green": 13, "blue": 14}
	for color, maxOccurrances := range maxNumberOfColors {
		matches := extractColor(line, color)
		for _, match := range matches {
			if matchInt, _ := strconv.Atoi(match); matchInt > maxOccurrances {
				return false
			}
		}
	}
	return true
}

func getMax(matches []string) int {
	max := -1
	for _, match := range matches {
		matchInt, _ := strconv.Atoi(match)
		if max == -1 || matchInt > max{
			max = matchInt
		}
	}
	return max
}

func partTwo(line string) int {
	results := map[string]int{"red": 1, "green": 1, "blue": 1}
	for color, _ := range results {
		matched := extractColor(line, color)
		results[color] = getMax(matched)
	}
	product := 1
	for _, v := range results {
		product *= v
	}
	return product
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	splitData := strings.Split(string(dat), "\n")
	result := 0
	for _, line := range splitData {
		// if isGameValid(line) {
		// 	result += i + 1
		// }
		result += partTwo(line)
	}
	fmt.Println(result)
}