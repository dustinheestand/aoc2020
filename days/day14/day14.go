package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var input []step

type step struct {
	ones      int64
	zeros     int64
	floatings int64
	vals      []tuple
}

type tuple struct {
	addr, val int64
}

func init() {
	file, err := os.Open("input/day14.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	s := step{}
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "mask") {
			input = append(input, s)
			s = step{}
			fs := strings.Fields(txt)
			for _, r := range fs[2] {
				s.ones <<= 1
				s.floatings <<= 1
				s.zeros <<= 1
				s.zeros |= 1
				switch r {
				case '0':
					s.zeros >>= 1
					s.zeros <<= 1
				case '1':
					s.ones |= 1
				case 'X':
					s.floatings |= 1
				}
			}
			continue
		}
		fs := strings.Fields(txt)
		reg := regexp.MustCompile("\\d+")
		addrStr := reg.Find([]byte(fs[0]))
		addr, err := strconv.ParseInt(string(addrStr), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		val, err := strconv.ParseInt(fs[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		s.vals = append(s.vals, tuple{addr, val})
	}
	input = append(input, s)
}

// Solve1 solves.
func Solve1() string {
	vals := make(map[int64]int64)
	for _, s := range input {
		for _, t := range s.vals {
			addr, v := t.addr, t.val
			v &= s.zeros
			v |= s.ones
			vals[addr] = v
		}
	}
	var res int64
	for _, v := range vals {
		res += v
	}
	return fmt.Sprint(res)
}

// Solve2 solves.
func Solve2() string {
	vals := make(map[int64]int64)
	for _, s := range input {
		for _, t := range s.vals {
			addr, v := t.addr, t.val
			f := s.floatings
			i := 0
			addrs := []int64{addr | s.ones}
			for f > 0 {
				if f&1 == 1 {
					addrs = float(addrs, i)
				}
				f >>= 1
				i++
			}
			for _, a := range addrs {
				vals[a] = v
			}
		}
	}
	var res int64
	for _, v := range vals {
		res += v
	}
	return fmt.Sprint(res)
}

func float(ns []int64, pos int) []int64 {
	var res []int64
	for _, n := range ns {
		res = append(res, n|(1<<pos))
		res = append(res, n & ^(1<<pos))
	}
	return res
}
