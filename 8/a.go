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

				newx1 := antennas[id][i].x + dx
				newy1 := antennas[id][i].y + dy
				if newx1 >= 0 && newx1 < width && newy1 >= 0 && newy1 < height {
					seen[fmt.Sprintf("%d.%d", newx1, newy1)] = true
				}

				newx2 := antennas[id][j].x - dx
				newy2 := antennas[id][j].y - dy
				if newx2 >= 0 && newx2 < width && newy2 >= 0 && newy2 < height {
					seen[fmt.Sprintf("%d.%d", newx2, newy2)] = true
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
