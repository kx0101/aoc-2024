package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	var numbers1, numbers2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])

		numbers1 = append(numbers1, num1)
		numbers2 = append(numbers2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file: ", err)
		return
	}

	sort.Ints(numbers1)
	sort.Ints(numbers2)

	sum := 0.0

	for i := 0; i < len(numbers1); i++ {
		sum += math.Abs(float64(numbers1[i]) - float64(numbers2[i]))
	}

	fmt.Println("sum: ", sum)

	wSum := 0

	for _, num1 := range numbers1 {
		count := 0

		for _, num2 := range numbers2 {
			if num1 == num2 {
				count++
			}
		}

		wSum += count * num1
	}

	fmt.Println("wSum: ", wSum)
}
