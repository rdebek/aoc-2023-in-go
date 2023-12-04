package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func partTwo(myNumbers []string, winningNumbers []string) int {
	numberOfWinningNumbers := 0
	for _, myNumber := range myNumbers {
		for _, winningNumber := range winningNumbers {
			if myNumber == winningNumber {
				numberOfWinningNumbers += 1
			}
		}
	}
	return numberOfWinningNumbers
}


func partOne(myNumbers []string, winningNumbers []string) int {
	numberOfWinningNumbers := 0
	for _, myNumber := range myNumbers {
		for _, winningNumber := range winningNumbers {
			if myNumber == winningNumber {
				numberOfWinningNumbers += 1
			}
		}
	}
	if numberOfWinningNumbers == 1 {
		return numberOfWinningNumbers
	} else {
		return int(math.Pow(2, float64(numberOfWinningNumbers - 1)))
	}
	return numberOfWinningNumbers
}


func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	splitData := strings.Split(string(dat), "\n")
	res := 0
	mapping := make(map[int]int, 0)
	for i, line := range splitData {
		res += 1
		usefulData := strings.Split(line, ":")[1]
		splitUsefulData := strings.Split(usefulData, "|")
		winningNumbers, myNumbers := strings.Split(strings.TrimSpace(strings.Join(strings.Fields(splitUsefulData[0]), " ")), " "), strings.Split(strings.TrimSpace(strings.Join(strings.Fields(splitUsefulData[1]), " ")), " ")
		// res += partOne(myNumbers, winningNumbers)
		if mapping[i] == 0 {
			mapping[i] = 1
		} else {
			mapping[i] += 1
		}
		for k := 0; k < mapping[i]; k++ {
			amountOfCards := partTwo(myNumbers, winningNumbers)
			res += amountOfCards
			for j := 1; j <= amountOfCards; j++ {
				mapping[i+j] += 1
			}
		}
	}
	fmt.Println(res)
}