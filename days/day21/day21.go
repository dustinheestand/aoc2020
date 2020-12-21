package day21

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var input []recipe

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
			t[f[:len(f)-o]] = true
		}
		input = append(input, r)
	}
}

// Solve1 solves.
func Solve1() string {
	allAllergens := make(map[string](map[string]bool))
	for _, r := range input {
		for a := range r.allergens {
			if allergens, ok := allAllergens[a]; !ok {
				allAllergens[a] = map[string]bool{}
				for i := range r.ingredients {
					allAllergens[a][i] = true
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

	maybeAllergens := make(map[string]bool)
	for _, is := range allAllergens {
		for i := range is {
			maybeAllergens[i] = true
		}
	}

	res := 0
	for _, r := range input {
		for i := range r.ingredients {
			if !maybeAllergens[i] {
				res++
			}
		}
	}
	return fmt.Sprint(res)
}

// Solve2 solves.
func Solve2() string {
	allAllergens := make(map[string](map[string]bool))
	for _, r := range input {
		for a := range r.allergens {
			if allergens, ok := allAllergens[a]; !ok {
				allAllergens[a] = map[string]bool{}
				for i := range r.ingredients {
					allAllergens[a][i] = true
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

	maybeAllergens := make(map[string]bool)
	for _, is := range allAllergens {
		for i := range is {
			maybeAllergens[i] = true
		}
	}

	seen := make(map[string]string)
	for len(seen) < 8 {
		for a, is := range allAllergens {
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

	allergens := sort.StringSlice{}
	for a := range allAllergens {
		allergens = append(allergens, a)
	}
	allergens.Sort()

	ings := []string{}
	for _, a := range allergens {
		for i := range allAllergens[a] {
			ings = append(ings, i)
		}
	}

	return fmt.Sprint(strings.Join(ings, ","))
}
