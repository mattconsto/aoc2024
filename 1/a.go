package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	pattern := regexp.MustCompile(`^\s*(\d+) +(\d+)\s*$`)
	scanner := bufio.NewScanner(os.Stdin)

	var lista []int
	var listb []int
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

		lista = append(lista, int(a))

		b, err := strconv.ParseInt(m[2], 10, 32)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing number:", err)
			os.Exit(1)
		}

		listb = append(listb, int(b))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	sort.Ints(lista)
	sort.Ints(listb)

	var total int
	for i := 0; i < len(lista); i++ {
		total += absDiffInt(lista[i], listb[i])
	}

	fmt.Println(total)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
