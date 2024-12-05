package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %s", err)
		return
	}

	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	var rules []string
	var updates []string

	i := 0
	for i < len(input) && strings.TrimSpace(input[i]) != "" {
		rules = append(rules, input[i])
		i++
	}

	i++
	for i < len(input) {
		updates = append(updates, input[i])
		i++
	}

	rulePairs := make([][2]int, len(rules))
	for i, rule := range rules {
		parts := strings.Split(rule, "|")

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		rulePairs[i] = [2]int{x, y}
	}

	sumMiddlePages := 0

	for _, update := range updates {
		pages := parseUpdateToInts(update)

		if !isValidOrder(pages, rulePairs) {
			sortedPages := topologicalSort(pages, rulePairs)
			middlePages := sortedPages[len(sortedPages)/2]

			sumMiddlePages += middlePages
		}
	}

	fmt.Printf("sum of middle pages: %d\n", sumMiddlePages)
}

func parseUpdateToInts(update string) []int {
	pageStrings := strings.Split(update, ",")
	pages := make([]int, len(pageStrings))

	for i, pageString := range pageStrings {
		page, _ := strconv.Atoi(pageString)
		pages[i] = page
	}

	return pages
}

func isValidOrder(pages []int, rules [][2]int) bool {
	for _, rule := range rules {
		indexX := indexOf(pages, rule[0])
		indexY := indexOf(pages, rule[1])

		if indexX != -1 && indexY != -1 && indexX > indexY {
			return false
		}
	}

	return true
}

func topologicalSort(pages []int, rules [][2]int) []int {
	graph := make(map[int][]int)
	inComingEdges := make(map[int]int)

	for _, page := range pages {
		graph[page] = []int{}
		inComingEdges[page] = 0
	}

	for _, rule := range rules {
		if contains(pages, rule[0]) && contains(pages, rule[1]) {
			graph[rule[0]] = append(graph[rule[0]], rule[1])
			inComingEdges[rule[1]]++
		}
	}

	queue := []int{}
	for node, edge := range inComingEdges {
		if edge == 0 {
			queue = append(queue, node)
		}
	}

	var sorted []int

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		sorted = append(sorted, curr)

		for _, neighbor := range graph[curr] {
			inComingEdges[neighbor]--

			if inComingEdges[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func indexOf(array []int, element int) int {
	for i, value := range array {
		if value == element {
			return i
		}
	}

	return -1
}

func contains(array []int, element int) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}

	return false
}
