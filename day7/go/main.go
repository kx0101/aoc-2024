package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Printf("error opening file: %s", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	totalResult := int64(0)

	for _, line := range lines {
		parts := strings.Split(line, ":")

		testValue, _ := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		numbersStrings := strings.Fields(strings.TrimSpace(parts[1]))

		numbers := make([]int64, len(numbersStrings))

		for i, ns := range numbersStrings {
			numbers[i], _ = strconv.ParseInt(ns, 10, 64)
		}

		if matchesValue(testValue, numbers) {
			totalResult += testValue
		}
	}

	fmt.Printf("result: %d\n", totalResult)
}

func matchesValue(target int64, numbers []int64) bool {
	return tryOperators(numbers, 1, numbers[0], target)
}

func tryOperators(numbers []int64, index int, currResult, target int64) bool {
	if index == len(numbers) {
		return currResult == target
	}

	nextNumber := numbers[index]

	if tryOperators(numbers, index+1, currResult+nextNumber, target) {
		return true
	}

	if tryOperators(numbers, index+1, currResult*nextNumber, target) {
		return true
	}

	combinedNumber := combineString(currResult, nextNumber)
	if tryOperators(numbers, index+1, combinedNumber, target) {
		return true
	}

	return false
}

func combineString(left, right int64) int64 {
	combinedNumber, _ := strconv.ParseInt(fmt.Sprintf("%d%d", left, right), 10, 64)
	return combinedNumber
}
