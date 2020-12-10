package day10

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

var input sort.IntSlice

func init() {
	file, err := os.Open("input/day10.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	input = append(input, 0)
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
	oneDiffs, threeDiffs := 0, 1
	last := 0
	for _, n := range input[1:] {
		if n-last == 1 {
			oneDiffs++
		} else if n-last == 3 {
			threeDiffs++
		}
		last = n
	}
	return strconv.Itoa(oneDiffs * threeDiffs)
}

// Solve2 solves.
func Solve2() string {
	ways := []int{0, 0, 1}
	for i, n := range input {
		if i == 0 {
			continue
		}
		waysToN := 0
		for j := 1; j <= 3 && i-j >= 0; j++ {
			if n-input[i-j] > 3 {
				break
			}
			waysToN += ways[3-j]
		}
		ways = append(ways, waysToN)
		ways = ways[1:]
	}
	return strconv.Itoa(ways[2])
}
