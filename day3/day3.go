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

func isDigit(character string) bool {
	_, err := strconv.Atoi(character)
	return err == nil
}

func isSymbol(character string) bool {
	return character == "*"
	// return !isDigit(character) && character != "."
}

func checkVerticalRow(aboveRow []string, startIndexX int, indexY int, alreadyVisitedCells [500][500]bool) ([]string, [500][500]bool) {
	var numbers = make([]string, 0)
	var wholeNumber string
	for i, character := range aboveRow[startIndexX: startIndexX + 3] {
		if isDigit(character) && alreadyVisitedCells[indexY][startIndexX + i] != true {
			wholeNumber, alreadyVisitedCells = getNumber(aboveRow, startIndexX + i, indexY, alreadyVisitedCells)
			numbers = append(numbers, wholeNumber)
		}
	}
	return numbers, alreadyVisitedCells
}


func checkHorizontalRowLeft(row []string, startIndexX int, indexY int, alreadyVisitedCells [500][500]bool) (string, [500][500]bool) {
	var wholeNumber string
	if isDigit(row[startIndexX]) {
		wholeNumber, alreadyVisitedCells = getLeftNumber(row, startIndexX, indexY, alreadyVisitedCells)
	}
	return wholeNumber, alreadyVisitedCells
}

func checkHorizontalRowRight(row []string, startIndexX int, indexY int, alreadyVisitedCells [500][500]bool) (string, [500][500]bool) {
	var wholeNumber string
	if isDigit(row[startIndexX]) {
		wholeNumber, alreadyVisitedCells = getRightNumber(row, startIndexX, indexY, alreadyVisitedCells)
	}
	return wholeNumber, alreadyVisitedCells
}

func getLeftNumber(row []string, index int, indexY int, alreadyVisitedCells [500][500]bool) (string, [500][500]bool){
	number := row[index]
	// check behind
	for i := index - 1; i >= 0; i-- {
		if alreadyVisitedCells[indexY][i] == true {
			return "0", alreadyVisitedCells
		}
		if isDigit(row[i]) {
			number = row[i] + number
			alreadyVisitedCells[indexY][i] = true
		} else {
			break
		}
	}
	return number, alreadyVisitedCells
}

func getRightNumber(row []string, index int, indexY int, alreadyVisitedCells [500][500]bool) (string, [500][500]bool){
	number := row[index]
	// check ahead
	for i := index + 1; i < len(row); i++ {
		if alreadyVisitedCells[indexY][i] == true {
			return "0", alreadyVisitedCells
		}
		if isDigit(row[i]) {
			number += row[i]
			alreadyVisitedCells[indexY][i] = true
		} else {
			break
		}
	}
	return number, alreadyVisitedCells
}



func getNumber(row []string, index int, indexY int, alreadyVisitedCells [500][500]bool) (string, [500][500]bool){
	number := row[index]
	// check behind
	for i := index - 1; i >= 0; i-- {
		if isDigit(row[i]) {
			number = row[i] + number
			alreadyVisitedCells[indexY][i] = true
		} else {
			break
		}
	}
	// check ahead
	for i := index + 1; i < len(row); i++ {
		if isDigit(row[i]) {
			number += row[i]
			alreadyVisitedCells[indexY][i] = true
		} else {
			break
		}
	}
	return number, alreadyVisitedCells

}


func partOne(splitLine [][]string) [][]string {
	numbers := make([][]string, 0)
	alreadyVisitedCells := [500][500]bool{}
	for i, line := range splitLine {
		for j, character := range line {
			if isSymbol(character) {
				numbersForGivenSymbol := make([]string, 0)
				aboveRow := make([]string, 0)
				aboveRow, alreadyVisitedCells = checkVerticalRow(splitLine[i - 1], j-1, i - 1, alreadyVisitedCells)
				for _, num := range aboveRow {
					numbersForGivenSymbol = append(numbersForGivenSymbol, num)
				}
				belowRow := make([]string, 0)
				belowRow, alreadyVisitedCells = checkVerticalRow(splitLine[i + 1], j-1, i + 1, alreadyVisitedCells)
				for _, num := range belowRow {
					numbersForGivenSymbol = append(numbersForGivenSymbol, num)
				}
				var leftNumber string
				leftNumber, alreadyVisitedCells = checkHorizontalRowLeft(splitLine[i], j-1, i, alreadyVisitedCells)
				numbersForGivenSymbol = append(numbersForGivenSymbol, leftNumber)

				var rightNumber string
				rightNumber, alreadyVisitedCells = checkHorizontalRowRight(splitLine[i], j+1, i, alreadyVisitedCells)
				numbersForGivenSymbol = append(numbersForGivenSymbol, rightNumber)
				
				numbersWithoutEmptyStrings := make([]string, 0)
				
				for _, num := range numbersForGivenSymbol {
					if num != "" {
						numbersWithoutEmptyStrings = append(numbersWithoutEmptyStrings, num)
					}
				}

				if len(numbersWithoutEmptyStrings) != 2 {
					numbersWithoutEmptyStrings = []string{}
				}
				numbers = append(numbers, numbersWithoutEmptyStrings)
			}
		}
	}
	return numbers
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	splitData := strings.Split(string(dat), "\n")
	arrayOfArrays := make([][]string, 0)
	for _, line := range splitData {
		line = line[:len(line)-1]
		arrayOfArrays = append(arrayOfArrays, strings.Split(line, ""))
	}
	result := 0
	numbers := partOne(arrayOfArrays)
	for _, numbersList := range numbers {
		if len(numbersList) > 0 {
		product := 1
		for _, number := range numbersList {
			numberInt, _ := strconv.Atoi(number)
			product *= numberInt
		}
		result += product
	}

	}
	fmt.Println(result)
	// result := 0
	// for _, number := range numbers {
	// 	convertedNumber, _ := strconv.Atoi(number) 
	// 	result += convertedNumber 
	// }
	// fmt.Println(result)
}