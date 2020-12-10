package day08

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type verb int

const (
	acc verb = iota
	jmp
	nop
)

var (
	verbs = map[string]verb{
		"acc": acc,
		"jmp": jmp,
		"nop": nop,
	}
)

func (i instruction) fix() instruction {
	if i.v == jmp {
		i.v = nop
	} else if i.v == nop {
		i.v = jmp
	}
	return i
}

type machine []instruction

func (m machine) run(pos, accum int, fix bool, seen map[int]bool) (res int, ok bool) {
	if _, ok := seen[pos]; ok {
		return accum, false
	}
	if pos == len(m) {
		return accum, true
	}
	seen[pos] = true
	inst := m[pos]
	if inst.v == acc {
		return m.run(pos+1, accum+inst.num, fix, seen)
	}
	instructions := []instruction{inst}
	if fix {
		instructions = append(instructions, inst.fix())
	}
	for idx, i := range instructions {
		copySeen := seen
		if fix {
			copySeen = make(map[int]bool, len(seen))
			for l := range seen {
				copySeen[l] = true
			}
		}
		if i.v == nop {
			res, ok = m.run(pos+1, accum, fix && idx == 0, copySeen)
		} else {
			res, ok = m.run(pos+i.num, accum, fix && idx == 0, copySeen)
		}
		if ok {
			return res, true
		}
	}
	return res, false
}

type instruction struct {
	v   verb
	num int
}

var input machine

func init() {
	file, err := os.Open("input/day08.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		fs := strings.Fields(txt)
		num, err := strconv.Atoi(fs[1])
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, instruction{verbs[fs[0]], num})
	}
}

// Solve1 solves.
func Solve1() string {
	res, _ := input.run(0, 0, false, map[int]bool{})
	return strconv.Itoa(res)
}

// Solve2 solves.
func Solve2() string {
	res, _ := input.run(0, 0, true, map[int]bool{})
	return strconv.Itoa(res)
}
