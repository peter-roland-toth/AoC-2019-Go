package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func check1(password string) bool {
	m := make(map[byte]int)
	m[password[0]] = 1
	for i := 1; i < len(password); i++ {
		if password[i] < password[i-1] {
			return false
		}
		m[password[i]]++
	}

	for _, v := range m {
		if v >= 2 {
			return true
		}
	}

	return false
}

func check2(password string) bool {
	m := make(map[byte]int)
	m[password[0]] = 1
	for i := 1; i < len(password); i++ {
		if password[i] < password[i-1] {
			return false
		}
		m[password[i]]++
	}

	for _, v := range m {
		if v == 2 {
			return true
		}
	}

	return false
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	input, _ := bufio.NewReader(file).ReadString('\n')
	str := strings.Split(input, "-")
	min, _ := strconv.Atoi(str[0])
	max, _ := strconv.Atoi(str[1])

	count1 := 0
	count2 := 0
	for p := min; p < max; p++ {
		pwd := strconv.Itoa(p)
		if check1(pwd) {
			count1++
		}
		if check2(pwd) {
			count2++
		}
	}
	
	fmt.Printf("Part 1: %d\n", count1)
	fmt.Printf("Part 2: %d\n", count2)
}
