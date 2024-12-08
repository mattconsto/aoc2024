package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	result := int64(0)
	for scanner.Scan() {
		line := scanner.Text()

		split1 := strings.Split(line, ": ")
		expected, err := strconv.ParseInt(split1[0], 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing number:", err)
			os.Exit(1)
		}

		split2 := strings.Split(split1[1], " ")
		var numbers []int64
		for i := 0; i < len(split2); i++ {
			n, err := strconv.ParseInt(split2[i], 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error parsing number:", err)
				os.Exit(1)
			}

			numbers = append(numbers, n)
		}

		for i := 0; i < (1 << len(numbers)); i++ {
			total := numbers[0]
			for j := 1; j < len(numbers); j++ {
				if i & (1 << j) != 0 {
					total += numbers[j]
				} else {
					total *= numbers[j]
				}
			}

			if total > expected {
				continue
			} else if total == expected {
				result += total
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	fmt.Println(result)
}