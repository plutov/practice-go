package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Point3D struct {
	X, Y, Z int
	Index   int
}

type Circuit struct {
	Points []int
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	points := make([]Point3D, len(lines))
	for i, line := range lines {
		coords := strings.Split(line, ",")
		if len(coords) != 3 {
			panic("Invalid input format")
		}
		p := Point3D{
			Index: i,
		}
		p.X, _ = strconv.Atoi(coords[0])
		p.Y, _ = strconv.Atoi(coords[1])
		p.Z, _ = strconv.Atoi(coords[2])
		points[i] = p
	}

	pairs := [][2]Point3D{}
	circuits := []Circuit{}
	for i := range points {
		circuits = append(circuits, Circuit{Points: []int{i}})

		for j := range points {
			if i == j {
				continue
			}
			if i > j {
				continue
			}
			pairs = append(pairs, [2]Point3D{points[i], points[j]})
		}
	}
	fmt.Println("Total pairs:", len(pairs))

	sort.SliceStable(pairs, func(i, j int) bool {
		return distance3D(pairs[i][0], pairs[i][1]) < distance3D(pairs[j][0], pairs[j][1])
	})

	for _, pair := range pairs {
		found := false
		for i, circuit := range circuits {
			if slices.Contains(circuit.Points, pair[0].Index) || slices.Contains(circuit.Points, pair[1].Index) {
				if !slices.Contains(circuit.Points, pair[1].Index) {
					circuits[i].Points = append(circuits[i].Points, pair[1].Index)
				}
				if !slices.Contains(circuit.Points, pair[0].Index) {
					circuits[i].Points = append(circuits[i].Points, pair[0].Index)
				}

				found = true

				break
			}
		}

		if !found {
			circuits = append(circuits, Circuit{Points: []int{pair[0].Index, pair[1].Index}})
		}

		circuits = mergeCircuits(circuits)

		if len(circuits) == 1 {
			fmt.Printf("Last pair: %d, %d\n", pair[0].Index, pair[1].Index)
			fmt.Printf("Res: %d\n", pair[0].X*pair[1].X)
			break
		}
	}

	sort.SliceStable(circuits, func(i, j int) bool {
		return len(circuits[i].Points) > len(circuits[j].Points)
	})
}

func distance3D(p1, p2 Point3D) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) +
		math.Pow(float64(p1.Y-p2.Y), 2) +
		math.Pow(float64(p1.Z-p2.Z), 2))
}

func mergeCircuits(circuits []Circuit) []Circuit {
	for {
		merged := false

		for i, c1 := range circuits {
			for j, c2 := range circuits {
				if i == j {
					continue
				}

				for _, p1 := range c1.Points {
					if slices.Contains(c2.Points, p1) {
						for _, p := range c2.Points {
							if !slices.Contains(circuits[i].Points, p) {
								circuits[i].Points = append(circuits[i].Points, p)
							}
						}
						circuits = append(circuits[:j], circuits[j+1:]...)
						merged = true
						break
					}
				}
				if merged {
					break
				}
			}
			if merged {
				break
			}
		}

		if !merged {
			break
		}
	}

	return circuits
}
