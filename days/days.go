package days

import (
	"github.com/dustinheestand/aoc2020/days/day1"
	"github.com/dustinheestand/aoc2020/days/day2"
	"github.com/dustinheestand/aoc2020/days/day3"
	"github.com/dustinheestand/aoc2020/days/day4"
	"github.com/dustinheestand/aoc2020/days/day5"
	"github.com/dustinheestand/aoc2020/days/day6"
)

// Day funcs.
type Day struct {
	Solve1 func() string
	Solve2 func() string
}

// Days is all the days.
var Days = map[int]Day{
	1: {day1.Solve1, day1.Solve2},
	2: {day2.Solve1, day2.Solve2},
	3: {day3.Solve1, day3.Solve2},
	4: {day4.Solve1, day4.Solve2},
	5: {day5.Solve1, day5.Solve2},
	6: {day6.Solve1, day6.Solve2},
}
