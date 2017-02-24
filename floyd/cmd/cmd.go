package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kennygrant/practice-go/floyd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please supply the number of rows you want as an argument\n")
		return
	}

	rows, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, row := range floyd.Triangle(rows) {
		for i := range row {
			fmt.Printf("%-2d ", row[i])
		}
		fmt.Printf("\n")
	}

}
