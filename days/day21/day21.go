package day21

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	input                   []recipe
	ingCount                map[string]int
	allergensToMatchingIngs map[string](map[string]bool)
	maybeAllergenicIngs     map[string]bool
)

type recipe struct {
	ingredients map[string]bool
	allergens   map[string]bool
}

func init() {
	file, err := os.Open("input/day21.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	ingCount = make(map[string]int)
	allergensToMatchingIngs = make(map[string](map[string]bool))
	maybeAllergenicIngs = make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		r := recipe{map[string]bool{}, map[string]bool{}}
		t := r.ingredients
		o := 0
		for _, f := range strings.Fields(txt) {
			if f == "(contains" {
				t = r.allergens
				o = 1
				continue
			}
			if o == 0 {
				ingCount[f]++
			}
			t[f[:len(f)-o]] = true
		}
		input = append(input, r)
	}

	for _, r := range input {
		for a := range r.allergens {
			if allergens, ok := allergensToMatchingIngs[a]; !ok {
				allergensToMatchingIngs[a] = map[string]bool{}
				for i := range r.ingredients {
					allergensToMatchingIngs[a][i] = true
				}
			} else {
				for a := range allergens {
					if !r.ingredients[a] {
						delete(allergens, a)
					}
				}
			}
		}
	}
	for _, is := range allergensToMatchingIngs {
		for i := range is {
			maybeAllergenicIngs[i] = true
		}
	}
}

// Solve1 solves.
func Solve1() string {
	res := 0
	for ing, count := range ingCount {
		if !maybeAllergenicIngs[ing] {
			res += count
		}
	}
	return fmt.Sprint(res)
}

// Solve2 solves.
func Solve2() string {
	seen := make(map[string]string)
	for len(seen) < 8 {
		for a, is := range allergensToMatchingIngs {
			for i := range is {
				if seen[i] != "" && seen[i] != a {
					delete(is, i)
				}
			}
			if len(is) == 1 {
				for i := range is {
					seen[i] = a
				}
			}
		}
	}

	var allergens sort.StringSlice
	for a := range allergensToMatchingIngs {
		allergens = append(allergens, a)
	}
	allergens.Sort()

	var ings []string
	for _, a := range allergens {
		for i := range allergensToMatchingIngs[a] {
			ings = append(ings, i)
		}
	}

	return fmt.Sprint(strings.Join(ings, ","))
}
