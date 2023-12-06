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

func partOne(seeds []string, mappings [][][]int) int{
	currentLowestLocation := -1
	for _, seed := range seeds {
		seed = strings.TrimSpace(seed)
		seed, _ := strconv.Atoi(seed)
		for _, mapping := range mappings {
			for _, innerMapping := range mapping {
				if seed >= innerMapping[1] && seed < innerMapping[1] + innerMapping[2] {
					seed = seed + (innerMapping[0] - innerMapping[1])
					break
				}
			}

		}
		if currentLowestLocation == -1 || seed < currentLowestLocation {
			currentLowestLocation = seed
		}
}
	return currentLowestLocation
}

func partTwo(seeds []string, mappings [][][]int) int{
	currentLowestLocation := -1
	for i := 0; i < len(seeds); i += 2 {
		initialSeed, _ := strconv.Atoi(strings.TrimSpace(seeds[i]))
		numberOfSeeds, _ := strconv.Atoi(strings.TrimSpace(seeds[i + 1]))
		for seed := initialSeed; seed < initialSeed + numberOfSeeds; seed++ {
			newSeed := seed
			for _, mapping := range mappings {
				for _, innerMapping := range mapping {
					if newSeed >= innerMapping[1] && newSeed < innerMapping[1] + innerMapping[2] {
						newSeed = newSeed + (innerMapping[0] - innerMapping[1])
						break
					}
				}
			}
			if currentLowestLocation == -1 || newSeed < currentLowestLocation {
				currentLowestLocation = newSeed
			}
		}
	}

	return currentLowestLocation
}


func makeMappingFromLine(line string) []int {
	listOfDigits := strings.Split(line, " ")
	destinationStart, _ := strconv.Atoi(listOfDigits[0])
	sourceStart, _ := strconv.Atoi(listOfDigits[1])
	numberOfDigits, _ := strconv.Atoi(listOfDigits[2])
	res := []int{destinationStart, sourceStart, numberOfDigits}
	return res
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	splitData := strings.Split(string(dat), "\n")
	seeds := make([]string, 0)
	mappings := make([][][]int, 0)
	for i, line := range splitData {
		if i == 0 {
			seeds = strings.Split(strings.Split(line, ":")[1], " ")[1:]
		} else if strings.Index(line, ":") > -1 {
			mapping := make([][]int, 0)
			for j := i + 1; j < len(splitData); j++ {
				if len(splitData[j]) == 1 || len(splitData[j]) == 0 {
					mappings = append(mappings, mapping)
					break
				}
				newMapping := makeMappingFromLine(splitData[j][:len(splitData[j]) - 1])
				mapping = append(mapping, newMapping)
			}
		}
	}
	// fmt.Println(partOne(seeds, mappings))
	fmt.Println(partTwo(seeds, mappings))
}
