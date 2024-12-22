package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type position struct {
	x int64
	y int64
	n int64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := map[int64][]int64{}
	var trailheads []position
	for y := int64(0); scanner.Scan(); y++ {
		line := scanner.Text()
		for x := int64(0); x < int64(len(line)); x++ {
			n, _ := strconv.ParseInt(string(line[x]), 10, 64)
			grid[y] = append(grid[y], n)

			if n == 0 {
				trailheads = append(trailheads, position{x, y, n})
			}
		}
	}

	total := 0
	for len(trailheads) > 0 {
		trailhead := trailheads[0]
		trailheads = trailheads[1:]

		var queue []position
		queue = append(queue, trailhead)

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.n == 9 {
				total++
			} else {
				if node.y > 0 && grid[node.y-1][node.x] == node.n+1 {
					queue = append(queue, position{node.x, node.y - 1, node.n + 1})
				}

				if node.y < int64(len(grid)-1) && grid[node.y+1][node.x] == node.n+1 {
					queue = append(queue, position{node.x, node.y + 1, node.n + 1})
				}

				if node.x > 0 && grid[node.y][node.x-1] == node.n+1 {
					queue = append(queue, position{node.x - 1, node.y, node.n + 1})
				}

				if node.x < int64(len(grid[node.y])-1) && grid[node.y][node.x+1] == node.n+1 {
					queue = append(queue, position{node.x + 1, node.y, node.n + 1})
				}
			}
		}
	}
	fmt.Println(total)
}
