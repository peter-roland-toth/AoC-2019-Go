package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

type vector struct {
	x int
	y int
	z int
}

type moon struct {
	position vector
	velocity vector
}

func (v *vector) add(u vector) {
	v.x += u.x
	v.y += u.y
	v.z += u.z
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// computes the greatest common divisor of two integers
func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	return a
}

// computes the least common multiple of the given integers
func lcm(a, b int, ints ...int) int {
	res := (a * b) / gcd(a,b)

	for _, i := range ints {
		res = lcm(res, i)
	}

	return res
}

func energy(v vector) int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func moonEnergy(m moon) int {
	return energy(m.position) * energy(m.velocity)
}

func applyGravity(m1, m2 *moon) {
	diffX, diffY, diffZ := 0, 0, 0
	if m1.position.x > m2.position.x {
		diffX = -1
	} else if m1.position.x < m2.position.x {
		diffX = 1
	}

	if m1.position.y > m2.position.y {
		diffY = -1
	} else if m1.position.y < m2.position.y {
		diffY = 1
	}

	if m1.position.z > m2.position.z {
		diffZ = -1
	} else if m1.position.z < m2.position.z {
		diffZ = 1
	}

	m1.velocity.add(vector{diffX, diffY, diffZ})
	m2.velocity.add(vector{-diffX, -diffY, -diffZ})
}

func applyVelocity(m *moon) {
	m.position.add(vector{m.velocity.x, m.velocity.y, m.velocity.z})
}

func runOneIteration(moons []moon) {
	for i1 := 0; i1 < len(moons); i1++ {
		for i2 := i1; i2 < len(moons); i2++ {
			moon1 := &moons[i1]
			moon2 := &moons[i2]
			if moon1 != moon2 {
				applyGravity(moon1, moon2)
			}
		}
	}

	for i := range moons {
		applyVelocity(&moons[i])
	}
}

func simulate(times int, moons []moon) {
	for i := 0; i < times; i++ {
		runOneIteration(moons)
	}
}

func findCycles(moons []moon) (int, int, int) {
	initialX := [4]int{moons[0].position.x, moons[1].position.x, moons[2].position.x, moons[3].position.x}
	initialY := [4]int{moons[0].position.y, moons[1].position.y, moons[2].position.y, moons[3].position.y}
	initialZ := [4]int{moons[0].position.z, moons[1].position.z, moons[2].position.z, moons[3].position.z}

	cycleX, cycleY, cycleZ := 0, 0, 0

	// since each dimension changes independently, it's enough to find the cycle time of each of them separately
	// the overall cycle time will be the LCM of the three dimensions' cycle times
	for i := 1; cycleX == 0 || cycleY == 0 || cycleZ == 0; i++ {
		runOneIteration(moons)

		newX := [4]int{moons[0].position.x, moons[1].position.x, moons[2].position.x, moons[3].position.x}
		newY := [4]int{moons[0].position.y, moons[1].position.y, moons[2].position.y, moons[3].position.y}
		newZ := [4]int{moons[0].position.z, moons[1].position.z, moons[2].position.z, moons[3].position.z}

		if cycleX == 0 && newX == initialX {
			velX := [4]int{moons[0].velocity.x, moons[1].velocity.x, moons[2].velocity.x, moons[3].velocity.x}
			if velX == [4]int{0, 0, 0, 0} {
				cycleX = i
			}
		}

		if cycleY == 0 && newY == initialY {
			velY := [4]int{moons[0].velocity.y, moons[1].velocity.y, moons[2].velocity.y, moons[3].velocity.y}
			if velY == [4]int{0, 0, 0, 0} {
				cycleY = i
			}
		}

		if cycleZ == 0 && newZ == initialZ {
			velZ := [4]int{moons[0].velocity.z, moons[1].velocity.z, moons[2].velocity.z, moons[3].velocity.z}
			if velZ == [4]int{0, 0, 0, 0} {
				cycleZ = i
			}
		}
	}

	return cycleX, cycleY, cycleZ
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()	

	scanner := bufio.NewScanner(file)
	moons := []moon{}

	for scanner.Scan() {
		i := strings.Split(strings.ReplaceAll(scanner.Text(), ">", ""), ",")
		x, _ := strconv.Atoi(strings.Split(i[0], "=")[1])
		y, _ := strconv.Atoi(strings.Split(i[1], "=")[1])
		z, _ := strconv.Atoi(strings.Split(i[2], "=")[1])
		m := moon{vector{x, y, z}, vector{0, 0, 0}}
		moons = append(moons, m)
	}

	moonCopy := make([]moon, len(moons))
	copy(moonCopy, moons)
	simulate(1000, moons)

	total := 0
	for _, m := range moons {
		total += moonEnergy(m)
	}
	fmt.Printf("Part 1: %d\n", total)

	x, y, z := findCycles(moonCopy)
	fmt.Printf("Part 2: %d\n", lcm(x, y, z))
}
