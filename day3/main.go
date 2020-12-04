package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var grid []string

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, s.Text())
	}

	// Part 1
	fmt.Println(hit(grid, 3, 1))

	// Part 2
	slopes := []struct {
		right, down int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	product := 1
	for _, s := range slopes {
		product *= hit(grid, s.right, s.down)
	}
	fmt.Println(product)
}

func hit(grid []string, right, down int) int {
	var (
		pos, count, lineNr = 0, 0, -1
		tree               = "#"
		next               = down
	)
	for _, line := range grid {
		lineNr++
		if lineNr != next {
			continue
		}
		next += down

		pos += right
		if pos >= len(line) {
			pos -= len(line)
		}
		if line[pos] == tree[0] {
			count++
		}
	}

	return count
}