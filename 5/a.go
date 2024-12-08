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

		if valid {
			total += pages[(len(pages) - 1) / 2]
		}
	}
	fmt.Println(total)
}
