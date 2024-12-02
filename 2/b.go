package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := 0
REPORT:
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")

	TRY:
		for try := 0; try < len(levels); try++ {
			lastSign := 0
			for i := 0; i < len(levels)-1; i++ {
				if i == try {
					continue
				}

				nextPos := i + 1
				if nextPos == try {
					// Avoid reading out of bounds
					if i+2 > len(levels)-1 {
						continue
					}
					nextPos++
				}

				a, _ := strconv.ParseInt(levels[i], 10, 32)
				b, _ := strconv.ParseInt(levels[nextPos], 10, 32)

				var diff = absDiffInt(int(a), int(b))
				if diff < 1 || diff > 3 {
					continue TRY
				}

				// Previous check ensure values differ
				sign := absSignInt(int(a), int(b))
				if lastSign == 0 {
					lastSign = sign
				} else if lastSign != sign {
					continue TRY
				}
			}

			total++
			continue REPORT
		}
	}

	fmt.Println(total)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}

func absSignInt(x, y int) int {
	if x-y < 0 {
		return -1
	} else if x-y > 0 {
		return 1
	} else {
		return 0
	}
}
