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

	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), " ")
		stones := map[int64]int64{}
		for i := 0; i < len(raw); i++ {
			number, _ := strconv.ParseInt(string(raw[i]), 10, 64)
			stones[number]++
		}

		for i := 1; i <= 75; i++ {
			newStones := map[int64]int64{}
			for stone, count := range stones {
				if stone == 0 {
					newStones[1] += count
				} else {
					numberString := strconv.FormatInt(stone, 10)
					if len(numberString)%2 == 0 {
						a, _ := strconv.ParseInt(string(numberString[:(len(numberString)/2)]), 10, 64)
						b, _ := strconv.ParseInt(string(numberString[(len(numberString)/2):]), 10, 64)
						newStones[a] += count
						newStones[b] += count
					} else {
						newStones[stone*2024] += count
					}
				}
			}
			stones = newStones
			total := int64(0)
			for _, count := range stones {
				total += count
			}
			fmt.Println(i, total)
		}
	}
}
