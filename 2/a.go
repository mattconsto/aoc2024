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

		lastSign := 0
		for i := 1; i < len(levels); i++ {
			a, _ := strconv.ParseInt(levels[i-1], 10, 32)
			b, _ := strconv.ParseInt(levels[i], 10, 32)

			var diff = absDiffInt(int(a), int(b))
			if diff < 1 || diff > 3 {
				continue REPORT
			}

			sign := absSignInt(int(a), int(b))
			if lastSign == 0 {
				lastSign = sign
			} else if lastSign != sign {
				continue REPORT
			}
		}

		total++
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
