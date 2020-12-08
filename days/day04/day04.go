package day04

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type passport map[string]string

var input []passport

func init() {
	file, err := os.Open("input/day04.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	p := passport{}
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			input = append(input, p)
			p = passport{}
			continue
		}
		fields := strings.Fields(txt)
		for _, f := range fields {
			tuple := strings.Split(f, ":")
			p[tuple[0]] = tuple[1]
		}
	}
	input = append(input, p)
}

// Solve1 solves.
func Solve1() string {
	var count int
	for _, p := range input {
		if valid(p) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func valid(p passport) bool {
	if _, ok := p["byr"]; !ok {
		return false
	}
	if _, ok := p["iyr"]; !ok {
		return false
	}
	if _, ok := p["eyr"]; !ok {
		return false
	}
	if _, ok := p["hgt"]; !ok {
		return false
	}
	if _, ok := p["hcl"]; !ok {
		return false
	}
	if _, ok := p["ecl"]; !ok {
		return false
	}
	if _, ok := p["pid"]; !ok {
		return false
	}
	return true
}

// Solve2 solves.
func Solve2() string {
	var count int
	for i, p := range input {
		if valid2(p) {
			if i == 22 {
				fmt.Println(p)
			}
			count++
		}
	}
	return strconv.Itoa(count)
}

var ecls = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func valid2(p passport) bool {
	byr := p["byr"]
	if yr, err := strconv.Atoi(byr); err != nil || yr < 1920 || yr > 2002 {
		return false
	}
	iyr := p["iyr"]
	if yr, err := strconv.Atoi(iyr); err != nil || yr < 2010 || yr > 2020 {
		return false
	}
	eyr := p["eyr"]
	if yr, err := strconv.Atoi(eyr); err != nil || yr < 2020 || yr > 2030 {
		return false
	}
	hgtReg := regexp.MustCompile("^(\\d+)(cm|in)$")
	hgt, ok := p["hgt"]
	match := hgtReg.Find([]byte(hgt))
	if match == nil {
		return false
	}
	fmt.Println(match)
	var num int
	for i, r := range hgt {
		if unicode.IsDigit(r) {
			d := int(hgt[i]) - 48
			num = num*10 + d
		} else if r == 'c' {
			if num > 193 || num < 150 || len(hgt) != i+2 || hgt[i+1] != 'm' {
				return false
			}
			break
		} else if r == 'i' {
			if num > 76 || num < 59 || len(hgt) != i+2 || hgt[i+1] != 'n' {
				return false
			}
			break
		} else {
			return false
		}
	}
	hcl, ok := p["hcl"]
	if !ok || hcl[0] != '#' || len(hcl) != 7 {
		return false
	}
	if _, err := hex.DecodeString(hcl[1:]); err != nil {
		return false
	}
	if ecl, ok := p["ecl"]; !ok || !ecls[ecl] {
		return false
	}
	pid, ok := p["pid"]
	if !ok || len(pid) != 9 {
		return false
	}
	if _, err := strconv.Atoi(pid); err != nil {
		return false
	}
	return true
}
