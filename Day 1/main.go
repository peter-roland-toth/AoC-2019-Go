package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	p1, p2 := 0, 0
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		p1 += fuel(i)
		p2 += allFuel(i)
	}

	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
}

func fuel(f int) int {
	return f / 3 - 2
}

func allFuel(f int) int {
	f = fuel(f)
	all := 0
	for f > 0 {
		all += f
		f = fuel(f)
	}

	return all
}
