package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	pattern := regexp.MustCompile(`^\s*(\d+) +(\d+)\s*$`)
	scanner := bufio.NewScanner(os.Stdin)

	var list []int
	var seen map[int]int = map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()

		m := pattern.FindStringSubmatch(line)
		if m == nil {
			fmt.Fprintln(os.Stderr, "Invalid input format: ", line)
			os.Exit(1)
		}

		a, err := strconv.ParseInt(m[1], 10, 32)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing number:", err)
			os.Exit(1)
		}

		list = append(list, int(a))

		b, err := strconv.ParseInt(m[2], 10, 32)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing number:", err)
			os.Exit(1)
		}

		if _, ok := seen[int(b)]; ok {
			seen[int(b)] = seen[int(b)] + 1
		} else {
			seen[int(b)] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	var total int
	for i := 0; i < len(list); i++ {
		total += list[i] * seen[int(list[i])]
	}

	fmt.Println(total)
}
