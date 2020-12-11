package day11

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var input []row

type row []int8

func init() {
	file, err := os.Open("input/day11.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	input = append(input, row{})
	for scanner.Scan() {
		txt := scanner.Text()
		r := row{-2}
		for _, rn := range txt {
			switch rn {
			case 'L':
				r = append(r, 0)
			case '.':
				r = append(r, -1)
			}
		}
		r = append(r, -2)
		input = append(input, r)
	}
	lenRow := len(input[1])
	emptyRow := row{}
	for i := 0; i < lenRow; i++ {
		emptyRow = append(emptyRow, -2)
	}
	input = append(input, emptyRow)
	input[0] = emptyRow
}

// Solve1 solves.
func Solve1() string {
	inputCopy := make([]row, len(input))
	for i, r := range input {
		rowCopy := make(row, len(r))
		copy(rowCopy, r)
		inputCopy[i] = rowCopy
	}
	changed := true
	for i := 0; changed; i++ {
		changed = false
		for rowNum, r := range inputCopy {
			for seatNum, s := range r {
				if s < 0 {
					continue
				}
				occNeighbors := int8(0)
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if y == 0 && x == 0 {
							continue
						}
						if inputCopy[rowNum+y][seatNum+x] < 0 {
							continue
						}
						if y < 0 || (y == 0 && x < 0) {
							occNeighbors += (inputCopy[rowNum+y][seatNum+x] >> 1) & 1
						} else {
							occNeighbors += inputCopy[rowNum+y][seatNum+x] & 1
						}
					}
				}
				inputCopy[rowNum][seatNum] = (s & 1) << 1
				if s&1 == 1 && occNeighbors < 4 {
					inputCopy[rowNum][seatNum]++
				}
				if s&1 == 0 && occNeighbors == 0 {
					inputCopy[rowNum][seatNum]++
				}
				if change := inputCopy[rowNum][seatNum]; change == 1 || change == 2 {
					changed = true
				}
			}
		}
	}
	res := 0
	for _, r := range inputCopy {
		for _, s := range r {
			if s > 0 && s&1 == 1 {
				res++
			}
		}
	}
	return strconv.Itoa(res)
}

// Solve2 solves.
func Solve2() string {
	inputCopy := make([]row, len(input))
	for i, r := range input {
		rowCopy := make(row, len(r))
		copy(rowCopy, r)
		inputCopy[i] = rowCopy
	}
	changed := true
	for i := 0; changed; i++ {
		changed = false
		for rowNum, r := range inputCopy {
			for seatNum, s := range r {
				if s < 0 {
					continue
				}
				occNeighbors := int8(0)
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if y == 0 && x == 0 {
							continue
						}
						xRay, yRay := x, y
						for {
							neighb := inputCopy[rowNum+yRay][seatNum+xRay]
							if neighb == -2 {
								break
							}
							if neighb == -1 {
								xRay += x
								yRay += y
								continue
							}
							if y < 0 || (y == 0 && x < 0) {
								occNeighbors += (inputCopy[rowNum+yRay][seatNum+xRay] >> 1) & 1
							} else {
								occNeighbors += inputCopy[rowNum+yRay][seatNum+xRay] & 1
							}
							break
						}
					}
				}
				inputCopy[rowNum][seatNum] = (s & 1) << 1
				if s&1 == 1 && occNeighbors < 5 {
					inputCopy[rowNum][seatNum]++
				}
				if s&1 == 0 && occNeighbors == 0 {
					inputCopy[rowNum][seatNum]++
				}
				if change := inputCopy[rowNum][seatNum]; change == 1 || change == 2 {
					changed = true
				}
			}
		}
	}
	res := 0
	for _, r := range inputCopy {
		for _, s := range r {
			if s > 0 && s&1 == 1 {
				res++
			}
		}
	}
	return strconv.Itoa(res)
}
