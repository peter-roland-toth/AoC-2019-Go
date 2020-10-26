package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	width := 25
	height := 6

	// reading the input as an array of bytes ("0" equals 48, "1" equals 49 and "2" equals 50)
	input, _ := bufio.NewReader(file).ReadBytes('\n')
	layers := [][]byte{}
	for i := 0; i < len(input); i += width * height {
		layers = append(layers, input[i:i+width*height])
	}

	minZeros, onesTimesTwos := width * height, 0

	for _, l := range layers {
		zeros, ones, twos := 0, 0, 0
		for j := 0; j < len(l); j++ {
			if l[j] == 48 {
				zeros++
			} else if l[j] == 49 {
				ones++
			} else if l[j] == 50 {
				twos++
			}
		}

		if zeros < minZeros {
			onesTimesTwos = ones * twos
			minZeros = zeros
		}
	}

	fmt.Printf("Part 1: %d\n", onesTimesTwos)
	fmt.Printf("Part 2:\n")

	finalPicture := layers[0]

	for i := 1; i < len(layers); i++ {
		for j := 0; j < width * height; j++ {
			if finalPicture[j] == 50 {
				finalPicture[j] = layers[i][j]
			}
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if finalPicture[i*width+j] == 48 {
				fmt.Print("  ")
			} else if finalPicture[i*width+j] == 49 {
				fmt.Print("[]")
			}
		}
		fmt.Print("\n")
	}
}
