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
		var stones []int64
		for i := 0; i < len(raw); i++ {
			number, _ := strconv.ParseInt(string(raw[i]), 10, 64)
			stones = append(stones, number)
		}

		for i := 1; i <= 25; i++ {
			length := len(stones)
			for j := 0; j < length; j++ {
				if stones[j] == 0 {
					stones[j] = 1
				} else {
					numberString := strconv.FormatInt(stones[j], 10)
					if len(numberString)%2 == 0 {
						a, _ := strconv.ParseInt(string(numberString[:(len(numberString)/2)]), 10, 64)
						b, _ := strconv.ParseInt(string(numberString[(len(numberString)/2):]), 10, 64)
						stones[j] = a
						stones = append(stones, b)
					} else {
						stones[j] = stones[j] * 2024
					}
				}
			}
			fmt.Println(i, len(stones))
		}
	}
}
