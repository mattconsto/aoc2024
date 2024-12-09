package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type file struct {
	id     int
	length int8
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		raw := scanner.Text()

		disk := list.New()
		for i := 0; i < len(raw); i++ {
			id := -1
			if i%2 == 0 {
				id = i / 2
			}

			length, _ := strconv.ParseInt(string(raw[i]), 10, 8)
			disk.PushBack(file{id, int8(length)})
		}

		endFile := disk.Back()
	NEXTFILE:
		for endFile != nil && endFile != disk.Front() {
			startSpace := disk.Front()
			for startSpace != nil && startSpace.Value.(file).id != -1 {
				startSpace = startSpace.Next()
				if startSpace == endFile {
					endFile = nextFile(endFile)
					continue NEXTFILE
				}
			}

			for {
				if startSpace.Value.(file).length >= endFile.Value.(file).length {
					inserted := disk.InsertBefore(endFile.Value, startSpace)
					if startSpace.Value.(file).length > endFile.Value.(file).length {
						// TODO merge spaces together!
						disk.InsertAfter(file{-1, startSpace.Value.(file).length - endFile.Value.(file).length}, inserted)
					}

					newEndFile := nextFile(endFile)
					disk.InsertBefore(file{-1, endFile.Value.(file).length}, endFile)
					disk.Remove(endFile)
					disk.Remove(startSpace)
					endFile = newEndFile
					continue NEXTFILE
				}

				for ok := true; ok; ok = startSpace != nil && startSpace.Value.(file).id != -1 {
					startSpace = startSpace.Next()
					if startSpace == endFile {
						endFile = nextFile(endFile)
						continue NEXTFILE
					}
				}
			}

			endFile = nextFile(endFile)
		}

		fmt.Println(checksumDisk(disk))
	}
}

func nextFile(mark *list.Element) *list.Element {
	mark = mark.Prev()
	for mark != nil && mark.Value.(file).id == -1 {
		mark = mark.Prev()
	}
	return mark
}

func checksumDisk(list *list.List) int {
	checksum := 0
	mark := list.Front()
	for i := 0; mark != nil; {
		fileId := mark.Value.(file).id
		if fileId != -1 {
			for j := 0; j < int(mark.Value.(file).length); j++ {
				checksum += (i + j) * fileId
			}
		}
		i += int(mark.Value.(file).length)
		mark = mark.Next()
	}
	return checksum
}
