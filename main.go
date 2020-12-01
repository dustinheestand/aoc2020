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
	fmt.Printf(
		"Day %d\nPart 1: %s\nPart 2: %s\n",
		*dFlag,
		days.Days[*dFlag].Solve1(),
		days.Days[*dFlag].Solve2(),
	)
}
