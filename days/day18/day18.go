package day18

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

var input []expr

type expr []rune

func init() {
	file, err := os.Open("input/day18.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		e := expr{}
		for _, r := range txt {
			if r == ' ' {
				continue
			}
			e = append(e, r)
		}
		input = append(input, e)
	}
}

type symbol interface {
	num() int
	rn() rune
}

type sRune rune

func (s sRune) rn() rune       { return rune(s) }
func (s sRune) num() int       { return -1 }
func (s sRune) String() string { return string(s) }

type sInt int

func (s sInt) rn() rune { return '!' }
func (s sInt) num() int { return int(s) }

// Solve1 solves.
func Solve1() string {
	var res int
	for _, e := range input {
		var stack []symbol
		for _, r := range e {
			if unicode.IsDigit(r) && (len(stack) == 0 || stack[len(stack)-1].rn() == '(') {
				stack = append(stack, sInt(int(r)-48))
			} else if r == '+' || r == '*' || r == '(' {
				stack = append(stack, sRune(r))
			} else if r == ')' {
				stack = append(stack[:len(stack)-2], stack[len(stack)-1])
			} else if unicode.IsDigit(r) {
				n := int(r) - 48
				op := stack[len(stack)-1]
				arg := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				if op.rn() == '+' {
					stack = append(stack, sInt(arg.num()+n))
				} else if op.rn() == '*' {
					stack = append(stack, sInt(arg.num()*n))
				}
			}
			for len(stack) >= 3 && stack[len(stack)-3].num() > 0 && stack[len(stack)-1].num() > 0 && stack[len(stack)-2].rn() != '!' {
				a, b, op := stack[len(stack)-3].num(), stack[len(stack)-1].num(), stack[len(stack)-2].rn()
				stack = stack[:len(stack)-3]
				if op == '+' {
					stack = append(stack, sInt(a+b))
				} else if op == '*' {
					stack = append(stack, sInt(a*b))
				}
			}
		}
		res += stack[0].num()
	}
	return fmt.Sprint(res)
}

// Solve2 solves.
func Solve2() string {
	var res int
	for _, e := range input {
		var stack []symbol
		for _, r := range e {
			if unicode.IsDigit(r) && (len(stack) == 0 || stack[len(stack)-1].rn() == '(' || stack[len(stack)-1].rn() == '*') {
				stack = append(stack, sInt(int(r)-48))
			} else if r == '+' || r == '*' || r == '(' {
				stack = append(stack, sRune(r))
			} else if r == ')' {
				for len(stack) >= 3 && stack[len(stack)-3].num() > 0 && stack[len(stack)-1].num() > 0 && stack[len(stack)-2].rn() == '*' {
					a, b := stack[len(stack)-3].num(), stack[len(stack)-1].num()
					stack = stack[:len(stack)-3]
					stack = append(stack, sInt(a*b))
				}
				stack = append(stack[:len(stack)-2], stack[len(stack)-1])
			} else if unicode.IsDigit(r) {
				n := int(r) - 48
				op := stack[len(stack)-1]
				if op.rn() == '+' {
					arg := stack[len(stack)-2]
					stack = stack[:len(stack)-2]
					stack = append(stack, sInt(arg.num()+n))
				}
			}
			for len(stack) >= 3 && stack[len(stack)-3].num() > 0 && stack[len(stack)-1].num() > 0 && stack[len(stack)-2].rn() == '+' {
				a, b := stack[len(stack)-3].num(), stack[len(stack)-1].num()
				stack = stack[:len(stack)-3]
				stack = append(stack, sInt(a+b))
			}
		}
		for len(stack) >= 3 && stack[len(stack)-3].num() > 0 && stack[len(stack)-1].num() > 0 && stack[len(stack)-2].rn() == '*' {
			a, b := stack[len(stack)-3].num(), stack[len(stack)-1].num()
			stack = stack[:len(stack)-3]
			stack = append(stack, sInt(a*b))
		}
		res += stack[0].num()
	}
	return fmt.Sprint(res)
}
