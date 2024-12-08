package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var grid []string
	x := 0
	y := 0

	{
		i := 0
		for scanner.Scan() {
			line := scanner.Text()
			grid = append(grid, line)
			index := strings.Index(line, "^")
			if index >= 0 {
				x = index
				y = i
			}
			i++
		}
	}

	direction := "up"
	var seen map[string]bool = map[string]bool{}
	for true {
		seen[fmt.Sprintf("%d.%d", y, x)] = true

		dy := 0
		dx := 0
		if direction == "up" {
			dy = -1
		} else if direction == "right" {
			dx = 1
		} else if direction == "down" {
			dy = 1
		} else if direction == "left" {
			dx = -1
		}

		if (y+dy) < 0 || (y+dy) >= len(grid) || (x+dx) < 0 || (x+dx) >= len(grid[y]) {
			break
		}

		if grid[y+dy][x+dx] == '#' {
			if direction == "up" {
				direction = "right"
			} else if direction == "right" {
				direction = "down"
			} else if direction == "down" {
				direction = "left"
			} else if direction == "left" {
				direction = "up"
			}
		} else {
			y += dy
			x += dx
		}
	}

	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if seen[fmt.Sprintf("%d.%d", i, j)] {
				total++
			}
		}
	}
	fmt.Println(total)
}
