package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func changeForDirection(direction byte) complex128 {
	change := complex(0, 0)
	switch direction {
	case byte('D'):
		change = complex(1, 0)
	case byte('U'):
		change = complex(-1, 0)
	case byte('R'):
		change = complex(0, 1)
	case byte('L'):
		change = complex(0, -1)
	}

	return change
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()	

	scanner := bufio.NewScanner(file)
	paths := [][]string{}

	for scanner.Scan() {
		paths = append(paths, strings.Split(scanner.Text(), ","))
	}

	// storing each visited point and how many steps it takes to get there
	visited := map[complex128]int{}

	position := complex(0, 0)
	visited[position] = 0
	steps := 0

	// going through the path of the first wire and marking every visited position
	// note that we don't only mark endpoints but every point on the line between two endpoints
	for _, p := range paths[0] {
		direction := p[0]
		distance, _ := strconv.Atoi(p[1:])
		change := changeForDirection(direction)		

		for d := 0; d < distance; d, steps = d+1, steps+1 {
			position += change
			visited[position] = steps
		}
	}

	position = complex(0, 0)
	closestDistance, minSteps := math.MaxInt32, math.MaxInt32

	steps = 0

	// walking through the second path now and for each intersection checking if it's the closest
	// to (0, 0)
	for _, p := range paths[1] {
		direction := p[0]
		distance, _ := strconv.Atoi(p[1:])
		change := changeForDirection(direction)		

		for d := 0; d < distance; d, steps = d+1, steps+1 {
			position += change

			s, ok := visited[position]
			if ok {
				dist := abs(int(real(position))) + abs(int(imag(position)))
				closestDistance = min(dist, closestDistance)
				minSteps = min(s + steps, minSteps)
			}
		}
	}

	fmt.Printf("Part 1: %d\n", closestDistance)
	fmt.Printf("Part 2: %d\n", minSteps)
}
