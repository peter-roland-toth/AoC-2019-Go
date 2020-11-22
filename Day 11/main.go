package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
	"time"
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


		// fmt.Println(op)

		if op == "99" {
			fmt.Println("hallo")
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

	inp := make(chan int)
	output := make(chan int)
	done := make(chan int)
	amp1 := newAmp(ints, inp, output, done)

	go amp1.run()

	// inp <- 1

	pos := complex(0, 0)
	dir := complex(0, -1)

	m := map[complex128]int{}
	m[pos] = 1

	for {

		select {
		case d := <-done:
			break
			fmt.Println("done", d)
			break
		default:
			// fmt.Println("run")
			currentColor, ok := m[pos]
			if !ok {
				currentColor = 0
			}
			// fmt.Println("inp1")
			amp1.input <- currentColor
			// fmt.Println("outp1")
			color, d := <- amp1.output
			if !d {
				break
			}
			// fmt.Println("outp2")
			turn, d := <- amp1.output
			if !d {
				break
			}
			// fmt.Println(color, turn)
			if turn == 0 {
				dir = dir * complex(0, -1)
			} else {
				dir = dir * complex(0, 1)
			}
			m[pos] = color
			// fmt.Println(pos, dir)
			pos = pos + dir
			// fmt.Println(m)

			// time.Sleep(10 * time.Microsecond)
		}
	}

	count := 0
	minX, minY, maxX, maxY := 0, 0, 0, 0
	for k, _ := range m {
		x := int(real(k))
		y := int(imag(k))
		if x < minX {
			minX = x
		} else if x > maxX {
			maxX = x
		}

		if y < minY {
			minY = y
		} else if y > maxY {
			maxY = y
		}
		count += 1
	}

	plot := [6][43]int{}

	fmt.Println(minX, maxX, minY, maxY, count)

	for k, v := range m {
		x := int(real(k))
		y := int(imag(k))

		plot[y-minY][x-minX] = v
	}

	for y := 0; y < len(plot); y++ {
		for x := 0; x < len(plot[0]); x++ {
			if plot[y][x] == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print("[]")
			}
		}

		fmt.Print("\n")
	}

	time.Sleep(500 * time.Millisecond)
}
