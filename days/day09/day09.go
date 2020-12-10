package day09

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

var input []int64

func init() {
	file, err := os.Open("input/day09.txt")
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
		input = append(input, int64(i))
	}
}

// Solve1 solves.
func Solve1() string {
	// How many of each sum we've seen.
	sums := make(map[int64]int)
	var sumsInclPos [](map[int64]bool)
	for i, n := range input {
		for j := 0; j < 25 && j < i; j++ {
			s := n + input[i+j-len(sumsInclPos)]
			sums[s]++
			sumsInclPos[j][s] = true
		}
		sumsInclPos = append(sumsInclPos, map[int64]bool{})
		if i <= 24 {
			continue
		}
		if sums[n] == 0 {
			return strconv.Itoa(int(n))
		}
		for obsoleteSum := range sumsInclPos[0] {
			sums[obsoleteSum]--
		}
		sumsInclPos = sumsInclPos[1:]
	}
	return ""
}

// Solve2 solves.
func Solve2() string {
	target := int64(177777905)
	var sum int64
	var i, start, end int
	for i < len(input) {
		sum += input[i]
		end = i + 1
		for sum > target {
			sum -= input[start]
			start++
		}
		if sum == target {
			var big, little int64 = -math.MaxInt64, math.MaxInt64
			for _, n := range input[start:end] {
				if n > big {
					big = n
				}
				if n < little {
					little = n
				}
			}
			return strconv.Itoa(int(big + little))
		}
		i++
	}
	return ""
}
