package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var input summary

// Fields are a slice of slice of ints (the ranges).
// Ticket is a one-length slice of ints (the numbers).
// Tickets is a slice (of tickets) of ints (the numbers).
type summary [][][]int

func init() {
	file, err := os.Open("input/day16.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var s [][]int
	for scanner.Scan() {
		txt := scanner.Bytes()
		if len(txt) == 0 {
			input = append(input, s)
			s = [][]int{}
			continue
		}
		r := regexp.MustCompile("\\d+")
		matches := r.FindAll(txt, -1)
		if len(matches) == 0 {
			continue
		}
		var ints []int
		for _, m := range matches {
			i, err := strconv.Atoi(string(m))
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, i)
		}
		s = append(s, ints)
	}
	input = append(input, s)
}

// Solve1 solves.
func Solve1() string {
	ranges, _, tickets := input[0], input[1][0], input[2]
	valid := map[int]bool{}
	for _, r := range ranges {
		for i := 0; i < 4; i += 2 {
			for j := r[i]; j <= r[i+1]; j++ {
				valid[j] = true
			}
		}
	}
	res := 0
	for _, t := range tickets {
		for _, n := range t {
			if !valid[n] {
				res += n
			}
		}
	}
	return fmt.Sprint(res)
}

// Solve2 solves.
func Solve2() string {
	ranges, myTicket, tickets := input[0], input[1][0], input[2]
	valid := map[int]bool{}
	for _, r := range ranges {
		for i := 0; i < 4; i += 2 {
			for j := r[i]; j <= r[i+1]; j++ {
				valid[j] = true
			}
		}
	}
	var validTickets [][]int
TicketsLoop:
	for _, t := range tickets {
		for _, n := range t {
			if !valid[n] {
				continue TicketsLoop
			}
		}
		validTickets = append(validTickets, t)
	}
	fields := make([]sort.IntSlice, len(validTickets[0]))
	for _, t := range validTickets {
		for i, f := range t {
			fields[i] = append(fields[i], f)
		}
	}
	for _, f := range fields {
		f.Sort()
	}
	matchUps := matchUp(fields)
	res := 1
	for _, m := range matchUps[:6] {
		res *= myTicket[m]
	}
	return fmt.Sprint(res)
}

type avails []avail
type avail struct {
	n   int
	pos sort.IntSlice
}

func (a avails) Less(i, j int) bool { return len(a[i].pos) < len(a[j].pos) }
func (a avails) Len() int           { return len(a) }
func (a avails) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func matchUp(fields []sort.IntSlice) []int {
	ranges := input[0]
	// The key represents the position in the list at the top.
	// The value represents the position on each ticket.
	fieldIDs := make([]int, 20)
	availPos := avails{}
	for i, r := range ranges {
		p := sort.IntSlice{}
		for j, f := range fields {
			if fits(r, f) {
				p = append(p, j)
			}
		}
		p.Sort()
		availPos = append(availPos, avail{i, p})
	}
	sort.Sort(availPos)

	tries := []int{0}
	for len(tries) <= len(ranges) {
		i := len(tries)
		lowest := tries[i-1]
		tries = tries[:i-1]
		taken := map[int]bool{}
		for p, t := range tries {
			taken[availPos[p].pos[t]] = true
		}
		for n := lowest; n < i; n++ {
			if !taken[availPos[i-1].pos[n]] {
				tries = append(tries, n)
				tries = append(tries, 0)
				break
			}
			if n == i-1 {
				for tries[len(tries)-1] == len(tries)-1 {
					tries = tries[:len(tries)-1]
				}
				tries[len(tries)-1]++
			}
		}
	}
	for i, t := range tries[:len(tries)-1] {
		fieldIDs[availPos[i].n] = availPos[i].pos[t]
	}
	return fieldIDs
}

func fits(ranges []int, values sort.IntSlice) bool {
	if len(values) == 0 {
		return true
	}
	if values[0] < ranges[0] {
		return false
	}
	if values[len(values)-1] > ranges[3] {
		return false
	}
	mid := len(values) / 2
	if values[mid] > ranges[1] && values[mid] < ranges[2] {
		return false
	}
	if values[mid] <= ranges[1] {
		return fits(ranges, values[mid+1:])
	}
	if values[mid] >= ranges[2] {
		return fits(ranges, values[:mid])
	}
	return false
}
