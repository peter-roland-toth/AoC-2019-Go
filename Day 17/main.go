package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

type amplifier struct {
	program []int // this is the "memory" of the amplifier, storing the program
	ip int // instruction pointer
	input chan int // input channel
	output chan int // output channel
	done chan int // channel indicating that the program has halted
	base int // relative base
}

func newAmp(program []int, input chan int, output chan int, done chan int) amplifier {
	cpy := make([]int, len(program))
	copy(cpy, program)
	memory := make([]int, 100000)
	cpy = append(cpy, memory...)
	return amplifier{cpy, 0, input, output, done, 0}
}

func (a *amplifier) run() {
	for true {
		x := 100000 + a.program[a.ip]
		s := strconv.Itoa(x)
		runes := []rune(s)
		op := string(runes[4:6])

		if op == "99" {
			a.done <- 1
			close(a.output)
			close(a.input)
			close(a.done)
			
			return
		}

		mode1 := string(runes[3])
		mode2 := string(runes[2])
		mode3 := string(runes[1])

		if op == "01" {
			param1 := a.program[a.ip+1]
			param2 := a.program[a.ip+2]
			param3 := a.program[a.ip+3]

			if mode1 == "0" {
				param1 = a.program[param1]
			} else if mode1 == "2" {
				param1 = a.program[param1+a.base]
			}
			
			if mode2 == "0" {
				param2 = a.program[param2]
			} else if mode2 == "2" {
				param2 = a.program[param2+a.base]
			}
			
			if mode3 == "2" {
				param3 = param3+a.base
			}

			a.program[param3] = param1 + param2
			a.ip += 4
		} else if op == "02" {
			param1 := a.program[a.ip+1]
			param2 := a.program[a.ip+2]
			param3 := a.program[a.ip+3]

			if mode1 == "0" {
				param1 = a.program[param1]
			} else if mode1 == "2" {
				param1 = a.program[param1+a.base]
			}
			
			if mode2 == "0" {
				param2 = a.program[param2]
			} else if mode2 == "2" {
				param2 = a.program[param2+a.base]
			}
			
			if mode3 == "2" {
				param3 = param3+a.base
			}

			a.program[param3] = param1 * param2
			a.ip += 4
		} else if op == "03" {
			param := a.program[a.ip+1]
			if mode1 == "2" {
				param = param+a.base
			}

			rec := <-a.input
			a.program[param] = rec
			a.ip += 2
		} else if op == "04" {
			param := a.program[a.ip+1]
			if mode1 == "0" {
				param = a.program[param]
			} else if mode1 == "2" {
				param = a.program[param+a.base]
			}
			send := param

			a.output <- send
			a.ip += 2
		} else if op == "05" {
			param1 := a.program[a.ip+1]
			if mode1 == "0" {
				param1 = a.program[param1]
			} else if mode1 == "2" {
				param1 = a.program[param1+a.base]
			}
			param2 := a.program[a.ip+2]
			if mode2 == "0" {
				param2 = a.program[param2]
			} else if mode2 == "2" {
				param2 = a.program[param2+a.base]
			}

			if param1 != 0 {
				a.ip = param2
			} else {
				a.ip += 3
			}
		} else if op == "06" {
			param1 := a.program[a.ip+1]
			if mode1 == "0" {
				param1 = a.program[param1]
			} else if mode1 == "2" {
				param1 = a.program[param1+a.base]
			}
			param2 := a.program[a.ip+2]
			if mode2 == "0" {
				param2 = a.program[param2]
			} else if mode2 == "2" {
				param2 = a.program[param2+a.base]
			}

			if param1 == 0 {
				a.ip = param2
			} else {
				a.ip += 3
			}
		} else if op == "07" {
			param1 := a.program[a.ip+1]
			param2 := a.program[a.ip+2]
			param3 := a.program[a.ip+3]

			if mode1 == "0" {
				param1 = a.program[param1]
			} else if mode1 == "2" {
				param1 = a.program[param1+a.base]
			}
			
			if mode2 == "0" {
				param2 = a.program[param2]
			} else if mode2 == "2" {
				param2 = a.program[param2+a.base]
			}
			
			if mode3 == "2" {
				param3 = param3+a.base
			}

			if param1 < param2 {
			 	a.program[param3] = 1
			} else {
			 	a.program[param3] = 0
			}
			a.ip += 4
		} else if op == "08" {

			param1 := a.program[a.ip+1]
			param2 := a.program[a.ip+2]
			param3 := a.program[a.ip+3]

			if mode1 == "0" {
				param1 = a.program[param1]
			} else if mode1 == "2" {
				param1 = a.program[param1+a.base]
			}
			
			if mode2 == "0" {
				param2 = a.program[param2]
			} else if mode2 == "2" {
				param2 = a.program[param2+a.base]
			}
			
			if mode3 == "2" {
				param3 = param3+a.base
			}

			if param1 == param2 {
			 	a.program[param3] = 1
			} else {
			 	a.program[param3] = 0
			}
			a.ip += 4
		} else if op == "09" {
			param := a.program[a.ip+1]
			if mode1 == "0" {
				param = a.program[param]
			} else if mode1 == "2" {
				param = a.program[param + a.base]
			}

			a.base += param
			a.ip += 2
		}
	}
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

	out := make(chan int)
	done := make(chan int)
	amp := newAmp(ints, make(chan int), out, done)

	go amp.run()

	row := 0
	col := 0
	m := map[[2]int]bool{}

	loop:
	for {
		select {
		case <-done:
			break loop
		case x := <-out:
			if x == 10 {
				row++
				col = 0
			} else {
				if x == 35 {
					m[[2]int{row, col}] = true
				}
				col++
			}
		}
	}

	total := 0
	for pos := range m {
		_, ok := m[[2]int{pos[0]-1, pos[1]}]
		if !ok {
			continue
		}

		_, ok = m[[2]int{pos[0]+1, pos[1]}]
		if !ok {
			continue
		}
		_, ok = m[[2]int{pos[0], pos[1]-1}]
		if !ok {
			continue
		}
		_, ok = m[[2]int{pos[0], pos[1]+1}]
		if !ok {
			continue
		}

		total += pos[0] * pos[1]
	}

	ints[0] = 2
	amp = newAmp(ints, make(chan int), make(chan int), make(chan int))

	instructions := []int{65, 44, 65, 44, 66, 44, 67, 44, 67, 44, 65, 44, 66, 44, 67, 44, 65, 44, 66, 10, 
		int('L'), 44, int('1'), int('2'), 44, int('L'), 44, int('1'), int('2'), 44, int('R'), 44, int('1'), int('2'), 10,
		int('L'), 44, int('8'), 44, int('L'), 44, int('8'), 44, int('R'), 44, int('1'), int('2'), 44, int('L'), 44, int('8'), 44, int('L'), 44, int('8'), 10,
		int('L'), 44, int('1'), int('0'), 44, int('R'), 44, int('8'), 44, int('R'), 44, int('1'), int('2'), 10,
		int('n'), 10, 10}

	index := 0
	finalScore := 0
	go amp.run()

	loop2:
	for {
		select {
		case <- amp.done:
			break loop2
		case amp.input <- instructions[index]:
			index++
		case out := <- amp.output:
			if out < 200 {
				fmt.Print(string(out))
			} else {
				finalScore = out
			}
		}
	}

	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", finalScore)
}
