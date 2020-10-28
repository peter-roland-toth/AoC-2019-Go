package main

import (
	"os"
	"bufio"
	"fmt"
	"math"
	"sort"
)

type point struct {
	X int
	Y int
}

func distance(p1 point, p2 point) float64 {
	dx, dy := float64(p2.X - p1.X), float64(p2.Y - p1.Y)
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func angle(p1 point, p2 point) float64 {
	dx, dy := float64(p2.X - p1.X), float64(p2.Y - p1.Y)
	return math.Atan2(dx, dy)
}

func determineBestAsteroid(asteroids []point) (int, point) {
	mostHits := 0
	bestPoint := point{}

	for _, ast1 := range asteroids {
		set := map[float64]bool{}

		for _, ast2 := range asteroids {
			if ast1.X != ast2.X || ast1.Y != ast2.Y {
				angle := angle(ast1, ast2)
				set[angle] = true
			}
		}

		if len(set) > mostHits {
			mostHits = len(set)
			bestPoint = ast1
		}
	}

	return mostHits, bestPoint
}

func determineVaporizationOrder(from point, asteroids []point) []point {
	vaporizedAsteroids := []point{from}
	visitedAsteroids := map[point]bool{}

	type AsteroidToVisit struct {
		Angle float64
		Point point
	}

	for len(vaporizedAsteroids) != len(asteroids) {
		possibleTargets := map[float64]point{}

		// determining the closest non-vaporized asteroid for each angle
		for _, asteroid := range asteroids {
			_, present := visitedAsteroids[asteroid]
			if !present && asteroid != from {
				angle := angle(from, asteroid)
				currentPoint, ok := possibleTargets[angle]
				if !ok || distance(currentPoint, from) > distance(asteroid, from) {
					possibleTargets[angle] = asteroid
				}
			}
		}

		// sorting the asteroids in clockwise order
		toVisit := []AsteroidToVisit{}

		for k, v := range possibleTargets {
			toVisit = append(toVisit, AsteroidToVisit{k, v})
		}

		sort.Slice(toVisit, func(i, j int) bool {
			return toVisit[i].Angle > toVisit[j].Angle
		})

		// marking the asteroids as vaporized
		for _, ast := range toVisit {
			vaporizedAsteroids = append(vaporizedAsteroids, ast.Point)
			visitedAsteroids[ast.Point] = true
		}
	}

	return vaporizedAsteroids
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()	

	scanner := bufio.NewScanner(file)
	m := []point{}

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x := 0; x < len(line); x++ {
			if string(line[x]) == "#" {
				m = append(m, point{x, y})
			}
		}
	}

	hits, bestAsteroid := determineBestAsteroid(m)
	vaporized := determineVaporizationOrder(bestAsteroid, m)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", hits, vaporized[200].X * 100 + vaporized[200].Y)
}
