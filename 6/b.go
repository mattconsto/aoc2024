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
	startx := 0
	starty := 0

	{
		i := 0
		for scanner.Scan() {
			line := scanner.Text()
			grid = append(grid, line)
			index := strings.Index(line, "^")
			if index >= 0 {
				startx = index
				starty = i
			}
			i++
		}
	}

	total := 0
	for oy := 0; oy < len(grid); oy++ {
		for ox := 0; ox < len(grid[oy]); ox++ {
			x := startx
			y := starty
			direction := "up"
			var seen map[string]bool = map[string]bool{}
			for true {
				if seen[fmt.Sprintf("%s.%d.%d", direction, y, x)] {
					total++
					break
				}

				seen[fmt.Sprintf("%s.%d.%d", direction, y, x)] = true

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

				if grid[y+dy][x+dx] == '#' || (y+dy == oy && x+dx == ox) {
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
		}
	}

	fmt.Println(total)
}
