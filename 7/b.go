package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lut := map[int]int64{}
	for i := 0; i < 32; i++ {
		lut[i] = powInt64(3, int64(i))
	}

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

		for i := int64(0); i < lut[len(numbers)]; i++ {
			total := numbers[0]
			for j := 1; j < len(numbers); j++ {
				remainder := i / lut[j] % 3
				if remainder == 0 {
					total += numbers[j]
				} else if remainder == 1 {
					total *= numbers[j]
				} else if remainder == 2 {
					concatinated := strconv.FormatInt(total, 10) + strconv.FormatInt(numbers[j], 10)
					converted, err := strconv.ParseInt(concatinated, 10, 64)
					if err != nil {
						fmt.Fprintln(os.Stderr, "Error parsing number:", err)
						os.Exit(1)
					}
					total = converted
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

func powInt64(x, y int64) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}
