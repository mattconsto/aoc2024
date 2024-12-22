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

type fence struct {
	x    int
	y    int
	side int
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
			var queue []position
			queue = append(queue, position{x, y})

			area := 0
			var fences map[string]fence = map[string]fence{}
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
					fences[fmt.Sprintf("%d.%d.t", p.y, p.x)] = fence{p.x, p.y, 't'}
				}

				if p.y < len(grid)-1 && grid[p.y][p.x] == grid[p.y+1][p.x] {
					queue = append(queue, position{p.x, p.y + 1})
				} else {
					fences[fmt.Sprintf("%d.%d.b", p.y, p.x)] = fence{p.x, p.y, 'b'}
				}

				if p.x > 0 && grid[p.y][p.x] == grid[p.y][p.x-1] {
					queue = append(queue, position{p.x - 1, p.y})
				} else {
					fences[fmt.Sprintf("%d.%d.l", p.y, p.x)] = fence{p.x, p.y, 'l'}
				}

				if p.x < len(grid[p.y])-1 && grid[p.y][p.x] == grid[p.y][p.x+1] {
					queue = append(queue, position{p.x + 1, p.y})
				} else {
					fences[fmt.Sprintf("%d.%d.r", p.y, p.x)] = fence{p.x, p.y, 'r'}
				}

				area++
			}

			sides := 0
			for len(fences) > 0 {
				var k string
				for k = range fences {
					break
				}
				f := fences[k]

				sides++
				if f.side == 't' || f.side == 'b' {
					for x := f.x - 1; x >= 0 && contains(fences, fmt.Sprintf("%d.%d.%c", f.y, x, f.side)); x-- {
						delete(fences, fmt.Sprintf("%d.%d.%c", f.y, x, f.side))
					}

					for x := f.x + 1; x < len(grid[f.y]) && contains(fences, fmt.Sprintf("%d.%d.%c", f.y, x, f.side)); x++ {
						delete(fences, fmt.Sprintf("%d.%d.%c", f.y, x, f.side))
					}
				} else {
					for y := f.y - 1; y >= 0 && contains(fences, fmt.Sprintf("%d.%d.%c", y, f.x, f.side)); y-- {
						delete(fences, fmt.Sprintf("%d.%d.%c", y, f.x, f.side))
					}

					for y := f.y + 1; y < len(grid) && contains(fences, fmt.Sprintf("%d.%d.%c", y, f.x, f.side)); y++ {
						delete(fences, fmt.Sprintf("%d.%d.%c", y, f.x, f.side))
					}
				}

				delete(fences, k)
			}

			total += area * sides
		}
	}

	fmt.Println(total)
}

func contains[M ~map[K]V, K comparable, V any](m M, k K) bool {
	_, ok := m[k]
	return ok
}
