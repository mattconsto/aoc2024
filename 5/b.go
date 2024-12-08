package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rules := map[int64][]int64{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		split := strings.Split(line, "|")
		a, _ := strconv.ParseInt(split[0], 10, 64)
		b, _ := strconv.ParseInt(split[1], 10, 64)

		rules[b] = append(rules[b], a)
	}

	total := int64(0)
	for scanner.Scan() {
		line := scanner.Text()

		raw := strings.Split(line, ",")
		var pages []int64
		for i := 0; i < len(raw); i++ {
			c, _ := strconv.ParseInt(raw[i], 10, 64)
			pages = append(pages, c)
		}

		valid := true
		NOTVALID: for i := 0; i < len(pages); i++ {
			if ruleset, ok := rules[pages[i]]; ok {
				for b := 0; b < len(ruleset); b++ {
					for j := i + 1; j < len(pages); j++ {
						if ruleset[b] == pages[j] {
							valid = false
							break NOTVALID
						}
					}
				}
			}
		}

		if !valid {
			remaining := make([]int64, len(pages))
			_ = copy(remaining, pages)

			var starting []int64
			var ending []int64

			for len(remaining) > 0 {
				seenA := map[int64]bool{}
				seenB := map[int64]bool{}

				for b, _ := range rules {
					if contains(remaining, b) {
						for i := 0; i < len(rules[b]); i++ {
							if contains(remaining, rules[b][i]) {
								seenA[rules[b][i]] = true
								seenB[b] = true
							}
						}
					}
				}

				for i := 0; i < len(remaining); i++ {
					if !seenA[remaining[i]] {
						ending = append(ending, remaining[i])
						remaining = remove(remaining, i)
						break
					} else if !seenB[remaining[i]] {
						starting = append(starting, remaining[i])
						remaining = remove(remaining, i)
						break
					}
				}
			}

			slices.Reverse(ending)
			combined := append(starting, ending...)
			total += combined[(len(combined) - 1) / 2]
		}
	}
	fmt.Println(total)
}

func contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
