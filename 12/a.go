package main

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var grid []string
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	var seen map[string]bool = map[string]bool{}
	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			area := 0
			perimeter := 0

			var queue []position
			queue = append(queue, position{x, y})
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]

				key := fmt.Sprintf("%d.%d", p.y, p.x)
				if seen[key] {
					continue
				}
				seen[key] = true

				if p.y > 0 && grid[p.y][p.x] == grid[p.y-1][p.x] {
					queue = append(queue, position{p.x, p.y - 1})
				} else {
					perimeter++
				}

				if p.y < len(grid)-1 && grid[p.y][p.x] == grid[p.y+1][p.x] {
					queue = append(queue, position{p.x, p.y + 1})
				} else {
					perimeter++
				}

				if p.x > 0 && grid[p.y][p.x] == grid[p.y][p.x-1] {
					queue = append(queue, position{p.x - 1, p.y})
				} else {
					perimeter++
				}

				if p.x < len(grid[p.y])-1 && grid[p.y][p.x] == grid[p.y][p.x+1] {
					queue = append(queue, position{p.x + 1, p.y})
				} else {
					perimeter++
				}

				area++
			}

			total += area * perimeter
		}
	}

	fmt.Println(total)
}
