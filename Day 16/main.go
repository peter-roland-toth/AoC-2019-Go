package main

import (
	"os"
	"bufio"
	"fmt"
	"bytes"
)

func multiplierForDigit(index, digit int) int {
	pattern := []int{0, 1, 0, -1}

	return pattern[((index + 1) % (digit * len(pattern)) / digit)]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func runPhase(signal []int) []int {
	newPhase := []int{}

	for digit := range signal {
		sum := 0
		for index, nr := range signal {
			sum += multiplierForDigit(index, digit + 1) * nr
		}

		newPhase = append(newPhase, abs(sum) % 10)
	}

	return newPhase
}

func intFromArray(arr []int, limit int) int {
	result := 0
	for count := 0; count < limit; count++ {
		result *= 10
		result += arr[count]
	}

	return result
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	i, _ := bufio.NewReader(file).ReadBytes('\n')
	i = bytes.TrimSpace(i)
	input := []int{}
	for j := 0; j < len(i); j++ {
		input = append(input, int(i[j] - '0'))
	}
	inputCopy := make([]int, len(input))
	copy(inputCopy, input)

	offset := intFromArray(input, 7)

	for count := 0; count < 100; count++ {
		input = runPhase(input)
	}

	fmt.Printf("Part 1: %d\n", intFromArray(input, 8))

	slice := make([]int, len(inputCopy) * 10000 - offset)
	for i := range slice {
		slice[i] = inputCopy[(offset + i) % len(inputCopy)]
	}

	for i := 0; i < 100; i++ {
		s := 0
		for j := len(slice) - 1; j >= 0; j-- {
			s += slice[j]
			slice[j] = abs(s) % 10
		}
	}

	fmt.Printf("Part 2: %d\n", intFromArray(slice, 8))
}
