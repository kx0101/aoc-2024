package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println(processInput(input))
}

func processInput(input string) int {
	mulPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	togglePattern := regexp.MustCompile(`(do\(\)|don't\(\))`)

	isMulEnabled := true
	sum := 0
	currPosition := 0

	for currPosition < len(input) {
		mulMatch := mulPattern.FindStringSubmatchIndex(input[currPosition:])
		toggleMatch := togglePattern.FindStringSubmatchIndex(input[currPosition:])

		if len(mulMatch) > 0 && (len(toggleMatch) == 0 || mulMatch[0] < toggleMatch[0]) {
			x, _ := strconv.Atoi(input[currPosition+mulMatch[2] : currPosition+mulMatch[3]])
			y, _ := strconv.Atoi(input[currPosition+mulMatch[4] : currPosition+mulMatch[5]])

			if isMulEnabled {
				sum += x * y
			}

			currPosition += mulMatch[1]
		} else if len(toggleMatch) > 0 {
			toggleValue := input[currPosition+toggleMatch[0] : currPosition+toggleMatch[1]]

			if toggleValue == "do()" {
				isMulEnabled = true
			} else if toggleValue == "don't()" {
				isMulEnabled = false
			}

			currPosition += toggleMatch[1]
		} else {
			currPosition++
		}
	}

	return sum
}
