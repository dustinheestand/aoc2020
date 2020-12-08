package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type password struct {
	min  int
	max  int
	c    rune
	word string
}

var input []password

func init() {
	file, err := os.Open("input/day02.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var p password
		fields := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return r == ':' || r == '-' || r == ' '
		})
		p.min, _ = strconv.Atoi(fields[0])
		p.max, _ = strconv.Atoi(fields[1])
		p.c, _ = utf8.DecodeRuneInString(fields[2])
		p.word = fields[3]
		input = append(input, p)
	}
}

// Solve1 solves.
func Solve1() string {
	var validCount int
	for _, p := range input {
		if p.isValid1() {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}

func (p password) isValid1() bool {
	var count int
	for _, c := range p.word {
		if c == p.c {
			count++
		}
		if count > p.max {
			return false
		}
	}
	if count < p.min {
		return false
	}
	return true
}

// Solve2 solves.
func Solve2() string {
	var validCount int
	for _, p := range input {
		if p.isValid2() {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}

func (p password) isValid2() bool {
	if isAtIdx(p.word, p.c, p.min-1) != isAtIdx(p.word, p.c, p.max-1) {
		return true
	}
	return false
}

func isAtIdx(word string, char rune, idx int) bool {
	if idx >= len(word) {
		return false
	}
	return word[idx] == byte(char)
}
