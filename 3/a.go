package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	scanner := bufio.NewScanner(os.Stdin)

	total := 0
	for scanner.Scan() {
		matches := pattern.FindAllStringSubmatch(scanner.Text(), -1)
		for i := 0; i < len(matches); i++ {
			a, _ := strconv.ParseInt(matches[i][1], 10, 32)
			b, _ := strconv.ParseInt(matches[i][2], 10, 32)
			total += int(a) * int(b)
		}
	}
	fmt.Println(total)
}
