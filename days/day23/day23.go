package day23

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var input sort.IntSlice

func init() {
	file, err := os.Open("input/day23.txt")
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
	return fmt.Sprint("")
}

// Solve2 solves.
func Solve2() string {
	return fmt.Sprint("")
}
