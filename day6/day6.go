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

func partOne(times []string, distances []string) int{
	results := make([]int, 0)
	for i:= 0; i < len(times); i++ {
		numberOfWaysToBeatTheRecord := 0
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		for j := 1; j < time; j ++ {
			remainingTime := time - j
			distanceTravelled := j * remainingTime
			if distanceTravelled > distance {
				numberOfWaysToBeatTheRecord += 1
			}
		}
		results = append(results, numberOfWaysToBeatTheRecord)
	}
	finalResult := 1
	for _, res := range results {
		finalResult *= res
	}
	return finalResult
}

func partTwo(time string, distance string) int{
	timeInt, _ := strconv.Atoi(time)
	distanceInt, _ := strconv.Atoi(distance)
	numberOfWaysToBeatTheRecord := 0
	for j := 1; j < timeInt; j ++ {
		remainingTime := timeInt - j
		distanceTravelled := j * remainingTime
		if distanceTravelled > distanceInt {
			numberOfWaysToBeatTheRecord += 1
		}
	}
	return numberOfWaysToBeatTheRecord
}


func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	splitData := strings.Split(string(dat), "\n")
	timesData := strings.Split(strings.Split(splitData[0], ":")[1], " ")[1:]
	for i := 0; i < len(timesData); i++ {
		timesData[i] = strings.TrimSpace(timesData[i])
	}
	distanceData := strings.Split(strings.Split(splitData[1], ":")[1], " ")[1:]
	// fmt.Println(partOne(timesData, distanceData))
	time := strings.Join(timesData, "")
	distance := strings.Join(distanceData, "")
	fmt.Println(partTwo(time, distance))
}
