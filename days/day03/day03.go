package day03

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type forest struct {
	width int
	trees [](map[int]bool)
}

var input forest

func init() {
	file, err := os.Open("input/day03.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		input.width = len(bytes)
		trees := make(map[int]bool)
		for i, b := range bytes {
			if b == '#' {
				trees[i] = true
			}
		}
		input.trees = append(input.trees, trees)
	}
}

// Solve1 solves.
func Solve1() string {
	var res, col int
	for _, row := range input.trees {
		if row[col%input.width] {
			res++
		}
		col += 3
	}
	return strconv.Itoa(res)
}

type slope struct {
	right, down int
}

// Solve2 solves.
func Solve2() string {
	slopes := []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var encounters []int
	for _, s := range slopes {
		var col, hits int
		for rowNum := 0; rowNum < len(input.trees); rowNum += s.down {
			if input.trees[rowNum][col%input.width] {
				hits++
			}
			col += s.right
		}
		encounters = append(encounters, hits)
	}
	res := 1
	for _, e := range encounters {
		res *= e
	}
	return strconv.Itoa(res)
}
