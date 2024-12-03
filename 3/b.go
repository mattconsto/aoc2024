package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	pattern := regexp.MustCompile(`(do|don't|mul)\((?:(\d+),(\d+))?\)`)
	scanner := bufio.NewScanner(os.Stdin)

	enabled := true
	total := 0
	for scanner.Scan() {
		matches := pattern.FindAllStringSubmatch(scanner.Text(), -1)
		for i := 0; i < len(matches); i++ {
			if matches[i][1] == "do" {
				enabled = true
			} else if matches[i][1] == "don't" {
				enabled = false
			} else if matches[i][1] == "mul" && enabled {
				a, _ := strconv.ParseInt(matches[i][2], 10, 32)
				b, _ := strconv.ParseInt(matches[i][3], 10, 32)
				total += int(a) * int(b)
			}
		}
	}
	fmt.Println(total)
}
