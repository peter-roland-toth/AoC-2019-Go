package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

func calculateTotalDepth(key string, m map[string][]string, depth int) int {
	children, ok := m[key]
	if !ok {
		return depth
	}

	totalDepth := 0
	for _, child := range children {
		totalDepth += calculateTotalDepth(child, m, depth + 1)
	}
	totalDepth += depth

	return totalDepth
}

func find(current string, key string, m map[string][]string, path []string) []string {
	if current == key {
		return path
	}

	children, _ := m[current]

	for _, c := range children {
		newPath := find(c, key, m, append(path, current))
		if len(newPath) > 0 {
			return newPath
		}
	}

	return []string{}
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()	

	scanner := bufio.NewScanner(file)
	m := make(map[string][]string)

	for scanner.Scan() {
		i := strings.Split(scanner.Text(), ")")
		_, ok := m[i[0]]
		if ok {
			m[i[0]] = append(m[i[0]], i[1])
		} else {
			m[i[0]] = []string{i[1]}
		}
	}

	fmt.Printf("Part 1: %d\n", calculateTotalDepth("COM", m, 0))

	path1 := find("COM", "YOU", m, []string{})
	path2 := find("COM", "SAN", m, []string{})

	// going through the two paths until they intersect
	i := 0
	for path1[i] == path2[i] {
		i++
	}

	// the result is the sum of the lenght of the uncommon parts in the two paths
	fmt.Printf("Part 2: %d\n", len(path1) + len(path2) - 2 * i)
}
