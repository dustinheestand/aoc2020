package day13

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

var input sort.IntSlice

func init() {
	file, err := os.Open("input/day13.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		i, err := strconv.Atoi(txt)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, i)
	}
	input.Sort()
}

// Solve1 solves.
func Solve1() string {
	return ""
}

// Solve2 solves.
func Solve2() string {
	return ""
}
