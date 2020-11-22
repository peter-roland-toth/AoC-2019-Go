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
			// fmt.Print("i")
			rec := <-a.input
			rec = <-a.input
			fmt.Println("GET", rec)
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

			// fmt.Print("o")
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

func display(tile [30][100]int, score int) {
	for y := 0; y < len(tile); y++ {
		for x := 0; x < len(tile[0]); x++ {
			if tile[y][x] == 0 {
				fmt.Print(" ")
			} else if tile[y][x] == 1 {
				fmt.Print("#")
			} else if tile[y][x] == 2 {
				fmt.Print("$")
			} else if tile[y][x] == 3 {
				fmt.Print("-")
			} else if tile[y][x] == 4 {
				fmt.Print("*")
			}
		}

		fmt.Print("\n")
	}

	fmt.Println("Part 2: ", score)
}

func runRobot(program []int) int {
	input := make(chan int)
	output := make(chan int)
	done := make(chan int)
	amp1 := newAmp(program, input, output, done)

	go amp1.run()

	m := []int{}
	blocks := 0

	tile := [30][100]int{}
	score := 0
	joystick := 0

	getInput := func() int {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Command ['a' for left, 'd' for right]: ")
		text, _ := reader.ReadString('\n')
		if text == "a\n" {
			joystick = -1
		} else if text == "d\n" {
			joystick = 1
		} else {
			joystick = 0
		}
		return joystick
	}

	out:
	for {
		select {
		case <-done:
			break out
		case o := <- amp1.output:
			m = append(m, o)
			if len(m) == 3 {
				if m[0] == -1 && m[1] == 0 {
					score = m[2]
				} else {
					tile[m[1]][m[0]] = m[2]
					if m[2] == 2 {
						blocks++
					}
				}
				m = nil
			}
		case amp1.input <- joystick:
			display(tile, score)
			amp1.input <- getInput()
		}
	}

	return blocks
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
	m := runRobot(ints)

	ints[0] = 2

	runRobot(ints)

	fmt.Printf("Part 1: %d\n", m)
}
