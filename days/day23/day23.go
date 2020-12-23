package day23

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	input1, input2 *node
	nodes1, nodes2 []*node
)

type node struct {
	val  int
	next *node
}

func init() {
	file, err := os.Open("input/day23.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	nodes1 = make([]*node, 10)
	nodes2 = make([]*node, 1e6+1)
	scanner := bufio.NewScanner(file)
	input1, input2 = &node{}, &node{}
	curr1, curr2 := input1, input2
	for scanner.Scan() {
		txt := scanner.Text()
		for _, c := range txt {
			if curr1.next != nil {
				curr1 = curr1.next
				curr2 = curr2.next
			}
			i, _ := strconv.Atoi(string(c))
			curr1.val, curr2.val = i, i
			nodes1[i], nodes2[i] = curr1, curr2
			curr1.next, curr2.next = &node{}, &node{}
		}
		for i := 10; i <= 1000000; i++ {
			curr2 = curr2.next
			nodes2[i] = curr2
			curr2.val = i
			curr2.next = &node{}
		}
	}
	curr1.next, curr2.next = input1, input2
}

// Solve1 solves.
func Solve1() string {
	solve(input1, nodes1, 9, 100)
	res := strings.Builder{}
	curr := nodes1[1].next
	for i := 0; i < 8; i++ {
		res.WriteString(fmt.Sprint(curr.val))
		curr = curr.next
	}
	return fmt.Sprint(res.String())
}

// Solve2 solves.
func Solve2() string {
	solve(input2, nodes2, 1e6, 1e7)
	return fmt.Sprint(nodes2[1].next.val * nodes2[1].next.next.val)
}

func solve(n *node, nodes []*node, numCups, numTurns int) {
	curr := n
	for i := 0; i < numTurns; i++ {
		firstPickedUp := curr.next
		var lastPickedUp *node
		// Find destination.
		n3map := make(map[int]bool, 3)
		dest := curr.val - 1
		if dest == 0 {
			dest = numCups
		}
		n3curr := curr.next
		for i := 0; i < 3; i++ {
			n3map[n3curr.val] = true
			n3curr = n3curr.next
			if i == 1 {
				lastPickedUp = n3curr
			}
			if i == 2 {
				curr.next = n3curr // Set the next of current to point to the node 4 after.
			}
		}
		for n3map[dest] {
			dest--
			if dest == 0 {
				dest = numCups
			}
		}
		destNode := nodes[dest]
		lastPickedUp.next, destNode.next = destNode.next, firstPickedUp
		curr = curr.next
	}
}
