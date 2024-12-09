package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		raw := scanner.Text()

		var digits []int8
		for i := 0; i < len(raw); i++ {
			digit, _ := strconv.ParseInt(string(raw[i]), 10, 8)
			digits = append(digits, int8(digit))
		}

		blockId := 0
		checksum := 0
		isBlock := true

		startDigit := 0
		startFile := 0

		endDigit := len(digits) - 1
		endFile := int(endDigit / 2)
	DONE:
		for true {
			for digits[startDigit] == 0 {
				startDigit++
				isBlock = !isBlock
				if isBlock {
					startFile++
				}
				if startDigit > endDigit {
					break DONE
				}
			}

			for digits[endDigit] == 0 {
				endDigit -= 2
				endFile--
				if endDigit <= startDigit {
					break DONE
				}
			}

			digits[startDigit]--
			if isBlock {
				checksum += blockId * startFile
			} else {
				checksum += blockId * endFile
				digits[endDigit]--
			}
			blockId++
		}
		fmt.Println(checksum)
	}
}
