package main

import (
	"flag"
	"fmt"

	"github.com/dustinheestand/aoc2020/days"
)

var (
	dFlag = flag.Int("day", 0, "day")
)

func main() {
	flag.Parse()
	if *dFlag != 0 {
		fmt.Printf(
			"Day %d\nPart 1: %s\nPart 2: %s\n",
			*dFlag,
			days.Days[*dFlag].Solve1(),
			days.Days[*dFlag].Solve2(),
		)
	}
	for i := 1; i <= len(days.Days); i++ {
		d := days.Days[i]
		fmt.Printf("Day %d  Part 1: %10v    Part 2: %10v\n", i, d.Solve1(), d.Solve2())
	}
}
