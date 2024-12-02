package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("error opening file: ", err)
		return
	}

	defer file.Close()

	reports := readReportsFromFile(file)

	safe := 0

	for _, report := range reports {
		if isSafe(report) || canBeSafeWithOneRemove(report) {
			safe++
		}
	}

	fmt.Println("safe: ", safe)
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	isIncreasing := levels[1] > levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}

		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			return false
		}
	}

	return true
}

func canBeSafeWithOneRemove(levels []int) bool {
	for i := range levels {
		modified := removeElement(levels, i)

		if isSafe(modified) {
			return true
		}
	}

	return false
}

func removeElement(levels []int, index int) []int {
	newSlice := append([]int{}, levels[:index]...)
	return append(newSlice, levels[index+1:]...)
}

func readReportsFromFile(file *os.File) [][]int {
	var reports [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		levelsStr := strings.Fields(line)

		var levels []int

		for _, s := range levelsStr {
			num, _ := strconv.Atoi(s)
			levels = append(levels, num)
		}

		reports = append(reports, levels)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file", err)
	}

	return reports
}
