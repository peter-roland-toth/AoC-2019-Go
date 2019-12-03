package main

import (
	"fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func mul(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

var m = map[int]func(int, int) int {
	1: add,
	2: mul,
}

func run(program []int, noun int, verb int) int {
	program[1] = noun
    program[2] = verb

    for i := 0; i < len(program); {
    	op := program[i]
    	
    	if op != 99 {
    		a := program[i+1]
    		b := program[i+2]
    		res := program[i+3]
    		program[res] = m[op](program[a], program[b])
    		i += 4
    	} else {
    		break
    	}
    }

    return program[0]
}

func main() {
	file, _ := os.Open("input")
    defer file.Close()

    input, _ := bufio.NewReader(file).ReadString('\n')
    str := strings.Split(input, ",")
    ints := []int{}
    for _, s := range str {
    	i, _ := strconv.Atoi(s) 
    	ints = append(ints, i)
    }  

    c := make([]int, len(ints))
    copy(c, ints)
    fmt.Printf("Part 1: %d\n", run(c, 12, 2))

    for n := 0; n < 100; n++ {
        for v:= 0; v < 100; v++ {
            copy(c, ints)
            if run(c, n, v) == 19690720 {
                fmt.Printf("Part 2: %d\n", n * 100 + v)
            }
        }
    }
 	 
}