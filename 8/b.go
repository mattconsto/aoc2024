package main

import (
	"bufio"
	"fmt"
	"os"
)

type antenna struct {
	x int
	y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	antennas := make(map[byte][]antenna)
	height := 0
	width := 0
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > width {
			width = len(line)
		}

		for x := 0; x < len(line); x++ {
			id := line[x]
			if id != '.' {
				antennas[id] = append(antennas[id], antenna{x, height})
			}
		}

		height++
	}

	seen := map[string]bool{}
	for id, _ := range antennas {
		for i := 0; i < len(antennas[id]); i++ {
			for j := i + 1; j < len(antennas[id]); j++ {
				dx := antennas[id][i].x - antennas[id][j].x
				dy := antennas[id][i].y - antennas[id][j].y

				for d := -1; d <= 1; d += 2 {
					for k := 0; true; k += d {
						newx := antennas[id][i].x + dx*k
						newy := antennas[id][i].y + dy*k

						if newx >= 0 && newx < width && newy >= 0 && newy < height {
							seen[fmt.Sprintf("%d.%d", newx, newy)] = true
						} else {
							break
						}
					}
				}
			}
		}
	}

	total := 0
	for _, _ = range seen {
		total++
	}
	fmt.Println(total)
}
