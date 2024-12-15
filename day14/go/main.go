package main

import (
	"bufio"
	"fmt"
	"os"
)

var max = P{101, 103}

type P struct{ x, y int }

type Robot struct {
	Pos, Vel P
}

func main() {
	robots := parseInput()

	for t := 0; t < 100; t++ {
		for i := 0; i < len(robots); i++ {
			robots[i].Pos.x = (robots[i].Pos.x + robots[i].Vel.x + max.x) % max.x
			robots[i].Pos.y = (robots[i].Pos.y + robots[i].Vel.y + max.y) % max.y
		}
	}

	q := [4]int{}
	for _, r := range robots {
		if r.Pos.x < max.x/2 && r.Pos.y < max.y/2 {
			q[0]++
		}

		if r.Pos.x < max.x/2 && r.Pos.y > max.y/2 {
			q[1]++
		}

		if r.Pos.x > max.x/2 && r.Pos.y < max.y/2 {
			q[2]++
		}

		if r.Pos.x > max.x/2 && r.Pos.y > max.y/2 {
			q[3]++
		}
	}

	fmt.Println(q[0] * q[1] * q[2] * q[3])
}

func parseInput() []Robot {
	file, err := os.Open("../input.txt")
	if err != nil {
		return nil
	}

	defer file.Close()

	robots := []Robot{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var x, y, vx, vy int

		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		robots = append(robots, Robot{P{x, y}, P{vx, vy}})
	}

	return robots
}
