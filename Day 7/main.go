package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func HeapPermutation(a []int, size int, result *[][]int) {
	if size == 1 {
		cp := make([]int, len(a))
		copy(cp, a)
		*result = append(*result, cp)
	}

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1, result)

		if size % 2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

type Amplifier struct {
	program []int // this is the "memory" of the amplifier, storing the program
	ip int // instruction pointer
	input chan int // input channel
	output chan int // output channel
	done chan int // channel indicating that the program has halted
}

func (a *Amplifier) run() {
	for true {
		x := 10000 + a.program[a.ip]
		s := strconv.Itoa(x)
		runes := []rune(s)
		op := string(runes[3:5])
		
		if op == "99" {
			a.done <- 1
			close(a.done)
			return
		}
		mode1 := string(runes[2])
		mode2 := string(runes[1])
		
		if op == "01" {
			param1 := a.program[a.ip+1]
			if mode1 == "0" {
				param1 = a.program[a.program[a.ip+1]]
			}
			param2 := a.program[a.ip+2]
			if mode2 == "0" {
				param2 = a.program[a.program[a.ip+2]]
			}
			param3 := a.program[a.ip+3]

			a.program[param3] = param1 + param2
			a.ip += 4
		} else if op == "02" {
			param1 := a.program[a.ip+1]
			if mode1 == "0" {
				param1 = a.program[a.program[a.ip+1]]
			}
			param2 := a.program[a.ip+2]
			if mode2 == "0" {
				param2 = a.program[a.program[a.ip+2]]
			}
			param3 := a.program[a.ip+3]
			a.program[param3] = param1 * param2
			a.ip += 4
		} else if op == "03" {
			rec := <-a.input
			a.program[a.program[a.ip+1]] = rec
			a.ip += 2
		} else if op == "04" {
			send := a.program[a.program[a.ip+1]]
			a.output <- send
			a.ip += 2
		} else if op == "05" {
			param1 := a.program[a.ip+1]
			if mode1 == "0" {
				param1 = a.program[a.program[a.ip+1]]
			}
			param2 := a.program[a.ip+2]
			if mode2 == "0" {
				param2 = a.program[a.program[a.ip+2]]
			}
			if param1 != 0 {
				a.ip = param2
			} else {
				a.ip += 3
			}
		} else if op == "06" {
			param1 := a.program[a.ip+1]
			if mode1 == "0" {
				param1 = a.program[a.program[a.ip+1]]
			}
			param2 := a.program[a.ip+2]
			if mode2 == "0" {
				param2 = a.program[a.program[a.ip+2]]
			}
			if param1 == 0 {
				a.ip = param2
			} else {
				a.ip += 3
			}
		} else if op == "07" {
			param1 := a.program[a.ip+1]
			if mode1 == "0" {
				param1 = a.program[a.program[a.ip+1]]
			}
			param2 := a.program[a.ip+2]
			if mode2 == "0" {
				param2 = a.program[a.program[a.ip+2]]
			}
			param3 := a.program[a.ip+3]
			if param1 < param2 {
			 a.program[param3] = 1
			} else {
			 a.program[param3] = 0
			}
			a.ip += 4
		} else if op == "08" {
			param1 := a.program[a.ip+1]
			if mode1 == "0" {
				param1 = a.program[a.program[a.ip+1]]
			}
			param2 := a.program[a.ip+2]
			if mode2 == "0" {
				param2 = a.program[a.program[a.ip+2]]
			}
			param3 := a.program[a.ip+3]
			if param1 == param2 {
			 a.program[param3] = 1
			} else {
			 a.program[param3] = 0
			}
			a.ip += 4
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
	
	phases := []int{0, 1, 2, 3, 4}
	permutations := [][]int{}
	HeapPermutation(phases, len(phases), &permutations)

	maxOutput := 0
	for _, per := range permutations {
		ampInput := 0
		for i := 0; i < 5; i++ {
			cpy := make([]int, len(ints))
			copy(cpy, ints)
			amp := Amplifier{cpy, 0, make(chan int), make(chan int), make(chan int)}
			
			// turning on the amp instance and sending the phase and the input
			go amp.run()
			amp.input <- per[i]
			amp.input <- ampInput

			// waiting for the output and using it as the next amp's input
			ampInput = <- amp.output
		}

		if ampInput > maxOutput {
			maxOutput = ampInput
		}
	}

	fmt.Printf("Part 1: %d\n", maxOutput)

	phases = []int{9, 8, 7, 6, 5}
	permutations = [][]int{}
	HeapPermutation(phases, len(phases), &permutations)

	maxOutput = 0

	for _, per := range permutations {
		channels := []chan int{}
		amps := []Amplifier{}

		// creating 5 channels which will be used by the amps to communicate
		for i := 0; i < 5; i++ {
			channels = append(channels, make(chan int))
		}
		for i := 0; i < 5; i++ {
			cpy := make([]int, len(ints))
			copy(cpy, ints)

			// assigning one channel as input and the next one as output
			// this way one amp's output becomes the next one's input
			amp := Amplifier{cpy, 0, channels[i], channels[(i+1)%5], make(chan int)}
			
			// turning the amp on and sending the phase number
			go amp.run()
			amps = append(amps, amp)
			amp.input <- per[i]
		}

		// sending 0 to the first channel (the first amp's input)
		channels[0] <- 0

		// waiting until the fourth amp is finished
		// this means that amp #5 will also finish and we can intercept its next output
		for range amps[3].done {
		}

		// intercepting the message sent to the first channel,
		// which is the output of amp #5
		ret := <- channels[0]

		if ret > maxOutput {
			maxOutput = ret
		}
	}

	fmt.Printf("Part 2: %d\n", maxOutput)
}

