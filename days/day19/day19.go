package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var input map[string]string
var input2 map[string]string
var matchStrings []string

type message interface {
	build(map[string]message) string
}

type sMessage string

func (s sMessage) build(map[string]message) string { return string(s) }

type rMessage struct {
	subs [][]string
}

func (r rMessage) build(inputs map[string]message) string {
	s := strings.Builder{}
	for _, sub := range r.subs {
		for _, m := range sub {
			s.WriteRune('(')
			s.WriteString(inputs[m].build(inputs))
			s.WriteRune(')')
		}
		s.WriteRune('|')
	}
	return s.String()[0 : len(s.String())-1]
}

type repMessage struct {
	reps []string
}

func (r repMessage) build(inputs map[string]message) string {
	s := strings.Builder{}
	for i := 1; i < 20; i++ {
		s.WriteRune('(')
		for _, rep := range r.reps {
			s.WriteRune('(')
			s.WriteString(inputs[rep].build(inputs))
			s.WriteRune(')')
			s.WriteString(fmt.Sprintf("{%v}", i))
		}
		s.WriteRune(')')
		s.WriteRune('|')
	}
	return s.String()[0 : len(s.String())-1]
}

func init() {
	file, err := os.Open("input/day19.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	input, input2 = make(map[string]string), make(map[string]string)
	scanner := bufio.NewScanner(file)
	defs := make(map[string]message)
	digits := regexp.MustCompile("\\d+")
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) == 0 {
			break
		}
		fs := strings.FieldsFunc(txt, func(r rune) bool { return r == ' ' })
		subs := [][]string{{}}
		key := fs[0][0 : len(fs[0])-1]
		for _, f := range fs[1:] {
			if f == "|" {
				subs = append(subs, []string{})
				continue
			}
			if !digits.Match([]byte(f)) {
				c := f[1:2]
				defs[key] = sMessage(c)
				break
			}
			subs[len(subs)-1] = append(subs[len(subs)-1], f)
		}
		if _, ok := defs[key]; !ok {
			defs[key] = rMessage{subs}
		}
	}
	for id, d := range defs {
		input[id] = d.build(defs)
	}
	defs["8"] = repMessage{[]string{"42"}}
	defs["11"] = repMessage{[]string{"42", "31"}}
	for id, d := range defs {
		input2[id] = d.build(defs)
	}
	for scanner.Scan() {
		matchStrings = append(matchStrings, scanner.Text())
	}
}

// Solve1 solves.
func Solve1() string {
	r := regexp.MustCompile("^" + input["0"] + "$")
	var res int
	for _, s := range matchStrings {
		if r.MatchString(s) {
			res++
		}
	}
	return fmt.Sprint(res)
}

// Solve2 solves.
func Solve2() string {
	r := regexp.MustCompile("^" + input2["0"] + "$")
	var res int
	for _, s := range matchStrings {
		if r.MatchString(s) {
			res++
		}
	}
	return fmt.Sprint(res)
}
