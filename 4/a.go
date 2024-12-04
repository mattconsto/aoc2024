package main

import (
	"bufio"
	"fmt"
	"os"
)

const WORD = "XMAS"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != WORD[0] {
				continue
			}

			for dy := -1; dy <= 1; dy++ {
			DIRECTION:
				for dx := -1; dx <= 1; dx++ {
					if dy == 0 && dx == 0 {
						continue DIRECTION
					}

					for i := 1; i < len(WORD); i++ {
						newy := y + dy*i
						newx := x + dx*i

						if newy < 0 ||
							newy >= len(grid) ||
							newx < 0 ||
							newx >= len(grid[newy]) ||
							grid[newy][newx] != WORD[i] {
							continue DIRECTION
						}
					}

					total++
				}
			}
		}
	}

	fmt.Println(total)
}
