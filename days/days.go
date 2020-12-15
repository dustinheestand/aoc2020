package days

import (
	"github.com/dustinheestand/aoc2020/days/day01"
	"github.com/dustinheestand/aoc2020/days/day02"
	"github.com/dustinheestand/aoc2020/days/day03"
	"github.com/dustinheestand/aoc2020/days/day04"
	"github.com/dustinheestand/aoc2020/days/day05"
	"github.com/dustinheestand/aoc2020/days/day06"
	"github.com/dustinheestand/aoc2020/days/day07"
	"github.com/dustinheestand/aoc2020/days/day08"
	"github.com/dustinheestand/aoc2020/days/day09"
	"github.com/dustinheestand/aoc2020/days/day10"
	"github.com/dustinheestand/aoc2020/days/day11"
	"github.com/dustinheestand/aoc2020/days/day12"
	"github.com/dustinheestand/aoc2020/days/day13"
	"github.com/dustinheestand/aoc2020/days/day14"
	"github.com/dustinheestand/aoc2020/days/day15"
	"github.com/dustinheestand/aoc2020/days/day16"
)

// Day funcs.
type Day struct {
	Solve1 func() string
	Solve2 func() string
}

// Days is all the days.
var Days = map[int]Day{
	1:  {day01.Solve1, day01.Solve2},
	2:  {day02.Solve1, day02.Solve2},
	3:  {day03.Solve1, day03.Solve2},
	4:  {day04.Solve1, day04.Solve2},
	5:  {day05.Solve1, day05.Solve2},
	6:  {day06.Solve1, day06.Solve2},
	7:  {day07.Solve1, day07.Solve2},
	8:  {day08.Solve1, day08.Solve2},
	9:  {day09.Solve1, day09.Solve2},
	10: {day10.Solve1, day10.Solve2},
	11: {day11.Solve1, day11.Solve2},
	12: {day12.Solve1, day12.Solve2},
	13: {day13.Solve1, day13.Solve2},
	14: {day14.Solve1, day14.Solve2},
	15: {day15.Solve1, day15.Solve2},
	16: {day16.Solve1, day16.Solve2},
}
