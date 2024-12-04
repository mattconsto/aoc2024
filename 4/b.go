package main

import (
	"bufio"
	"fmt"
	"os"
)

const WORD = "MAS"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	total := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != WORD[1] {
				continue
			}

			found := 0
			for dy := -1; dy <= 1; dy++ {
			DIRECTION:
				for dx := -1; dx <= 1; dx++ {
					if dy == 0 || dx == 0 {
						continue DIRECTION
					}

					for i := -1; i <= 1; i++ {
						newy := y + dy*i
						newx := x + dx*i
						if newy < 0 ||
							newy >= len(grid) ||
							newx < 0 ||
							newx >= len(grid[newy]) ||
							grid[newy][newx] != WORD[i+1] {
							continue DIRECTION
						}
					}

					found++
				}
			}

			if found == 2 {
				total++
			}
		}
	}

	fmt.Println(total)
}
