package day25

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input []int

func init() {
	file, err := os.Open("input/day25.txt")
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
}

// Solve1 solves.
func Solve1() string {
	loopSizes := [2]int{}
	sn := 7
	for i, key := range input {
		try := 1
		size := 0
		for try != key {
			try *= sn
			try %= 20201227
			size++
		}
		loopSizes[i] = size
	}
	res := int64(1)
	for i := 0; i < loopSizes[1]; i++ {
		res *= int64(input[0])
		res %= 20201227
	}
	return fmt.Sprint(res)
}

// Solve2 solves.
func Solve2() string {
	return fmt.Sprint("")
}
