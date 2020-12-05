package day5

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type seat struct {
	row, col int
}

type seats []seat

var input seats

func (ss seats) Len() int      { return len(ss) }
func (ss seats) Swap(i, j int) { ss[i], ss[j] = ss[j], ss[i] }
func (ss seats) Less(i, j int) bool {
	a, b := ss[i], ss[j]
	if a.row < b.row {
		return true
	}
	if a.row == b.row && a.col < b.col {
		return true
	}
	return false
}

func init() {
	file, err := os.Open("input/day5.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		rowStr := txt[:7]
		colStr := txt[7:]
		rowStr = strings.Replace(rowStr, "F", "0", -1)
		rowStr = strings.Replace(rowStr, "B", "1", -1)
		colStr = strings.Replace(colStr, "L", "0", -1)
		colStr = strings.Replace(colStr, "R", "1", -1)
		row, _ := strconv.ParseInt(rowStr, 2, 32)
		col, _ := strconv.ParseInt(colStr, 2, 32)
		input = append(input, seat{int(row), int(col)})
	}
	sort.Sort(input)
}

// Solve1 solves.
func Solve1() string {
	var max int
	for _, s := range input {
		idx := s.row*8 + s.col
		if max < idx {
			max = idx
		}
	}
	return strconv.Itoa(max)
}

// Solve2 solves.
func Solve2() string {
	lastIdx := -100
	for _, s := range input {
		idx := s.row*8 + s.col
		if idx-lastIdx == 2 {
			return strconv.Itoa(idx - 1)
		}
		lastIdx = idx
	}
	return ""
}
