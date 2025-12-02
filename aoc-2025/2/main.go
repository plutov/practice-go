package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	pairs := strings.Split(string(file), ",")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		pair = strings.Trim(pair, "\n")
		if pair == "" {
			continue
		}

		ids := strings.Split(pair, "-")
		lowestId, err := strconv.Atoi(ids[0])
		if err != nil {
			panic(err)
		}
		highestId, err := strconv.Atoi(ids[1])
		if err != nil {
			panic(err)
		}

		for i := lowestId; i <= highestId; i++ {
			idStr := strconv.Itoa(i)

			for partsCount := len(idStr); partsCount >= 2; partsCount-- {
				if len(idStr)%partsCount == 0 {
					partLen := len(idStr) / partsCount
					allEqual := true
					firstPart := idStr[:partLen]
					for j := 0; j < partsCount; j++ {
						if idStr[j*partLen:(j+1)*partLen] != firstPart {
							allEqual = false
							break
						}
					}

					if allEqual {
						sum += i
						break
					}
				}
			}
		}
	}

	fmt.Println(sum)
}
