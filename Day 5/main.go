package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func run(program []int, input int) int {
	ip := 0
	output := input
	for true {
		// adding 10000 to make sure the instruction doesn't change but it's always 4 characters long
		x := 10000 + program[ip]
		s := strconv.Itoa(x)
		runes := []rune(s)
	    op := string(runes[3:5])
	    if op == "99" {
	    	return output
	    } else {
	    	mode1 := string(runes[2])
		    mode2 := string(runes[1])
		    
		    if op == "01" {
		    	param1 := program[ip+1]
			    if mode1 == "0" {
			    	param1 = program[program[ip+1]]
			    }
			    param2 := program[ip+2]
			    if mode2 == "0" {
			    	param2 = program[program[ip+2]]
			    }
			    param3 := program[ip+3]

		    	program[param3] = param1 + param2
				ip += 4
		    } else if op == "02" {
		    	param1 := program[ip+1]
			    if mode1 == "0" {
			    	param1 = program[program[ip+1]]
			    }
			    param2 := program[ip+2]
			    if mode2 == "0" {
			    	param2 = program[program[ip+2]]
			    }
			    param3 := program[ip+3]
		    	program[param3] = param1 * param2
				ip += 4
		    } else if op == "03" {
		    	program[program[ip+1]] = output
				ip += 2
		    } else if op == "04" {
		    	output = program[program[ip+1]]
				ip += 2
		    } else if op == "05" {
		    	param1 := program[ip+1]
			    if mode1 == "0" {
			    	param1 = program[program[ip+1]]
			    }
			    param2 := program[ip+2]
			    if mode2 == "0" {
			    	param2 = program[program[ip+2]]
			    }
		    	if param1 != 0 {
					ip = param2
		    	} else {
					ip += 3
		    	}
		    } else if op == "06" {
		    	param1 := program[ip+1]
			    if mode1 == "0" {
			    	param1 = program[program[ip+1]]
			    }
			    param2 := program[ip+2]
			    if mode2 == "0" {
			    	param2 = program[program[ip+2]]
			    }
		    	if param1 == 0 {
					ip = param2
		    	} else {
					ip += 3
		    	}
		    } else if op == "07" {
		    	param1 := program[ip+1]
			    if mode1 == "0" {
			    	param1 = program[program[ip+1]]
			    }
			    param2 := program[ip+2]
			    if mode2 == "0" {
			    	param2 = program[program[ip+2]]
			    }
			    param3 := program[ip+3]
		    	if param1 < param2 {
					program[param3] = 1
		    	} else {
					program[param3] = 0
		    	}
		    	ip += 4
		    } else if op == "08" {
		    	param1 := program[ip+1]
			    if mode1 == "0" {
			    	param1 = program[program[ip+1]]
			    }
			    param2 := program[ip+2]
			    if mode2 == "0" {
			    	param2 = program[program[ip+2]]
			    }
			    param3 := program[ip+3]
		    	if param1 == param2 {
					program[param3] = 1
		    	} else {
					program[param3] = 0
		    	}
		    	ip += 4
		    }
	    }
	}

	return -1
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
	fmt.Printf("Part 1: %d\n", run(c, 1))
	copy(c, ints)
	fmt.Printf("Part 2: %d\n", run(c, 5))
}
