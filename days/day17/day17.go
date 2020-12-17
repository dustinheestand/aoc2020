package day17

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	input cube
	dim   int
)

// Fields are a slice of slice of ints (the ranges).
// Ticket is a one-length slice of ints (the numbers).
// Tickets is a slice (of tickets) of ints (the numbers).
type cube map[coord]int

type coord struct {
	x, y, z int
}

func init() {
	file, err := os.Open("input/day17.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	input = cube{}
	i := 0
	for scanner.Scan() {
		txt := scanner.Bytes()
		for j, r := range txt {
			if r == '#' {
				input[coord{i, j, 0}] = 1
			}
		}
		i++
	}
	dim = 8 // Width and height of the initial space.
}

// Solve1 solves.
func Solve1() string {
	cu := make(cube, len(input))
	for c := range input {
		cu[c] = 1
	}
	for cycle := 1; cycle <= 6; cycle++ {
		for i := -cycle; i < dim+cycle; i++ {
			for j := -cycle; j < dim+cycle; j++ {
				for k := -cycle; k < 1+cycle; k++ {
					co := coord{i, j, k}
					cu[co] *= 10
					alive := aliveLast(cu[co], cycle)
					adj := numAdj(co, cu, cycle)
					if (alive && adj == 2) || adj == 3 {
						cu[co] += cycle + 1
					}
				}
			}
		}
	}
	res := 0
	for _, v := range cu {
		if v%10 != 0 {
			res++
		}
	}
	return fmt.Sprint(res)
}

func numAdj(co coord, cu cube, cy int) int {
	x, y, z := co.x, co.y, co.z
	res := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				val := cu[coord{x + i, y + j, z + k}]
				if aliveLast(val, cy) {
					res++
				}
			}
		}
	}
	return res
}

func aliveLast(val, cy int) bool {
	return val%10 == cy || (val/10)%10 == cy
}

type coord4 struct {
	x, y, z, w int
}

func numAdj4(co coord4, cu map[coord4]int, cy int) int {
	x, y, z, w := co.x, co.y, co.z, co.w
	res := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for h := -1; h <= 1; h++ {
					if i == 0 && j == 0 && k == 0 && h == 0 {
						continue
					}
					val := cu[coord4{x + i, y + j, z + k, w + h}]
					if aliveLast(val, cy) {
						res++
					}
				}
			}
		}
	}
	return res
}

// Solve2 solves.
func Solve2() string {
	cu := make(map[coord4]int, len(input))
	for c := range input {
		cu[coord4{c.x, c.y, c.z, 0}] = 1
	}
	for cycle := 1; cycle <= 6; cycle++ {
		for i := -cycle; i < dim+cycle; i++ {
			for j := -cycle; j < dim+cycle; j++ {
				for k := -cycle; k < 1+cycle; k++ {
					for h := -cycle; h < 1+cycle; h++ {
						co := coord4{i, j, k, h}
						cu[co] *= 10
						alive := aliveLast(cu[co], cycle)
						adj := numAdj4(co, cu, cycle)
						if (alive && adj == 2) || adj == 3 {
							cu[co] += cycle + 1
						}
					}
				}
			}
		}
	}
	res := 0
	for _, v := range cu {
		if v%10 != 0 {
			res++
		}
	}
	return fmt.Sprint(res)
}
