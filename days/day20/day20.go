package day20

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	tilesPerRow  = 12
	rowsPerImage = 12 * 8
)

var input map[string]tile

type tile [10][10]bool

func (t tile) String() string {
	sb := strings.Builder{}
	for _, r := range t {
		for _, c := range r {
			if c {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func init() {
	file, err := os.Open("input/day20.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	input = map[string]tile{}
	t := tile{}
	var key string
	row := 0
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) == 0 {
			input[key] = t
			t = tile{}
			row = 0
			continue
		}
		if strings.HasPrefix(txt, "Tile") {
			key = txt[5 : len(txt)-1]
			continue
		}
		for c, r := range txt {
			if r == '#' {
				t[row][c] = true
			}
		}
		row++
	}
	input[key] = t
}

// Solve1 solves.
func Solve1() string {
	sides := make(map[[10]bool][]string)
	for id, t := range input {
		rotated := t.rotate90()
		for _, s := range [][10]bool{
			t[0],
			t[9],
			reverse(t[0]),
			reverse(t[9]),
			rotated[0],
			rotated[9],
			reverse(rotated[0]),
			reverse(rotated[9]),
		} {
			sides[s] = append(sides[s], id)
		}
	}
	// Find a corner
	edge := make(map[string]int)
	for _, ts := range sides {
		if len(ts) == 1 {
			edge[ts[0]]++
		}
	}
	res := 1
	for k, e := range edge {
		if e == 4 {
			val, err := strconv.Atoi(k)
			if err != nil {
				return ""
			}
			res *= val
		}
	}

	return fmt.Sprint(res)
}

func (t tile) rightSide() [10]bool {
	rotated := t.rotate90()
	return rotated[9]
}

func (t tile) leftSide() [10]bool {
	rotated := t.rotate90()
	return rotated[0]
}

func reverse(arr [10]bool) [10]bool {
	res := [10]bool{}
	for i, b := range arr {
		res[9-i] = b
	}
	return res
}

func (t tile) rotate90() tile {
	res := tile{}
	for i, r := range t {
		for j, b := range r {
			res[j][9-i] = b
		}
	}
	return res
}

func (t tile) flipHoriz() tile {
	res := tile{}
	for i, r := range t {
		for j, b := range r {
			res[i][9-j] = b
		}
	}
	return res
}

func (t tile) permutations() []tile {
	return []tile{t, t.rotate90(), t.rotate90().rotate90(), t.rotate90().rotate90().rotate90(),
		t.flipHoriz(), t.flipHoriz().rotate90(), t.flipHoriz().rotate90().rotate90(), t.flipHoriz().rotate90().rotate90().rotate90(),
	}
}

type image [rowsPerImage][rowsPerImage]bool

func (img image) String() string {
	sb := strings.Builder{}
	for _, r := range img {
		for _, b := range r {
			if b {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (img image) rotate90() image {
	res := image{}
	for i, r := range img {
		for j, b := range r {
			res[j][rowsPerImage-1-i] = b
		}
	}
	return res
}

func (img image) flipHoriz() image {
	res := image{}
	for i, r := range img {
		for j, b := range r {
			res[i][rowsPerImage-1-j] = b
		}
	}
	return res
}

func (img image) permutations() []image {
	return []image{img, img.rotate90(), img.rotate90().rotate90(), img.rotate90().rotate90().rotate90(),
		img.flipHoriz(), img.flipHoriz().rotate90(), img.flipHoriz().rotate90().rotate90(), img.flipHoriz().rotate90().rotate90().rotate90(),
	}
}

// Solve2 solves.
func Solve2() string {
	sides := make(map[[10]bool][]string)
	for id, t := range input {
		rotated := t.rotate90()
		for _, s := range [][10]bool{
			t[0],
			t[9],
			reverse(t[0]),
			reverse(t[9]),
			rotated[0],
			rotated[9],
			reverse(rotated[0]),
			reverse(rotated[9]),
		} {
			sides[s] = append(sides[s], id)
		}
	}
	allTiles := [tilesPerRow][tilesPerRow]tile{}
	placed := map[string]bool{}
	// Find a corner
	edge := make(map[string]int)
	for _, ts := range sides {
		if len(ts) == 1 {
			edge[ts[0]]++
		}
	}
	for key, ts := range edge {
		if ts == 4 {
			for _, t := range input[key].permutations() {
				if len(sides[t[0]]) == 1 && len(sides[t.leftSide()]) == 1 {
					allTiles[0][0] = t
					break
				}
			}
			placed[key] = true
			break
		}
	}

	for i := 0; i < tilesPerRow; i++ {
		for j := 0; j < tilesPerRow; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				leftMatch := allTiles[i][j-1].rightSide()
				for _, k := range sides[leftMatch] {
					if placed[k] {
						continue
					}
					for _, t := range input[k].permutations() {
						if t.leftSide() == leftMatch && len(sides[t[0]]) == 1 {
							allTiles[i][j] = t
						}
					}
					placed[k] = true
				}
				continue
			}
			if j == 0 {
				topMatch := allTiles[i-1][j][9]
				for _, k := range sides[topMatch] {
					if placed[k] {
						continue
					}
					for _, t := range input[k].permutations() {
						if t[0] == topMatch && len(sides[t.leftSide()]) == 1 {
							allTiles[i][j] = t
						}
					}
					placed[k] = true
				}
				continue
			}
			leftMatch := allTiles[i][j-1].rightSide()
			topMatch := allTiles[i-1][j][9]
			matchesTop := map[string]bool{}
			newTile := tile{}
			for _, k := range sides[topMatch] {
				if placed[k] {
					continue
				}
				matchesTop[k] = true
			}
			for _, k := range sides[leftMatch] {
				if matchesTop[k] {
					newTile = input[k]
					placed[k] = true
				}
			}
			for _, t := range newTile.permutations() {
				if t.leftSide() == leftMatch && t[0] == topMatch {
					allTiles[i][j] = t
				}
			}
		}
	}

	img := image{}
	for i, rowOfTiles := range allTiles {
		for j, t := range rowOfTiles {
			for k, row := range t[1:9] {
				for l, b := range row[1:9] {
					img[i*8+k][j*8+l] = b
				}
			}
		}
	}

	monsterHashes := map[coord]bool{}
	middleReg := regexp.MustCompile("a.{4}aa.{4}aa.{4}aaa")
	for _, perm := range img.permutations() {
		for i, row := range perm {
			if i == 0 || i == rowsPerImage-1 {
				continue // Skip the first and last rows.
			}
			rowStr := strings.Builder{}
			for _, b := range row {
				if b {
					rowStr.WriteRune('a')
				} else {
					rowStr.WriteRune('-')
				}
			}
			matches := middleReg.FindAllIndex([]byte(rowStr.String()), -1)
		matchLoop:
			for _, m := range matches {
				if !perm[i-1][m[0]+18] {
					continue
				}
				nextRow := perm[i+1]
				for _, offset := range []int{1, 4, 7, 10, 13, 16} {
					if !nextRow[m[0]+offset] {
						continue matchLoop
					}
				}
				for _, c := range monsterCoords(i, m[0]) {
					monsterHashes[c] = true
				}
			}
		}
		if len(monsterHashes) != 0 {
			break
		}
	}

	totalTrue := 0
	for _, r := range img {
		for _, b := range r {
			if b {
				totalTrue++
			}
		}
	}
	return fmt.Sprint(totalTrue - len(monsterHashes))
}

type coord struct {
	a, b int
}

func monsterCoords(a, b int) []coord {
	return []coord{
		{a - 1, b + 18},
		{a, b},
		{a, b + 5},
		{a, b + 6},
		{a, b + 11},
		{a, b + 12},
		{a, b + 17},
		{a, b + 18},
		{a, b + 19},
		{a + 1, b + 1},
		{a + 1, b + 4},
		{a + 1, b + 7},
		{a + 1, b + 10},
		{a + 1, b + 13},
		{a + 1, b + 16},
	}
}
