package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputText string

func main() {
	fmt.Printf("Magic id: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic id: %d\n", run2(inputText))
}

type present struct {
	id     int
	area   int
	layout []string
}
type region struct {
	id       int
	sizeX    int
	sizeY    int
	area     int
	presents []int
}

func run1(input string) int {
	presents, regions := parse(input)

	count := 0

	for _, r := range regions {
		tpc, tpa := sumPresentsAndAreas(presents, r.presents)
		if tpa > r.area {
			continue
		}
		if tpc*9 <= r.area {
			count++
			continue
		}
		log.Printf("region %d: tpa: %d region area: %d", r.id, tpa, r.area)
	}

	return count
}

func sumPresentsAndAreas(presents map[int]present, presentCounts []int) (totalPresents, totalArea int) {
	for id, howmany := range presentCounts {
		totalArea += howmany * presents[id].area
		totalPresents += howmany
	}
	return totalPresents, totalArea
}

func run2(input string) int {
	return 0
}

func parse(input string) (map[int]present, []region) {
	segments := SplitByEmptyNewlineToSlices(input)

	presentChunks := segments[:6]
	regionChunk := segments[6]

	presents := make(map[int]present)

	for _, pc := range presentChunks {
		p := present{}
		p.id, _ = strconv.Atoi(strings.TrimRight(pc[0], ":"))
		p.layout = pc[1:]
		p.area = strings.Count(pc[1], "#") + strings.Count(pc[2], "#") + strings.Count(pc[3], "#")
		presents[p.id] = p
	}

	regions := []region{}
	for i, rc := range regionChunk {
		r := region{id: i}
		rcc := strings.Split(rc, ":")
		fmt.Sscanf(rcc[0], "%dx%d", &r.sizeX, &r.sizeY)
		r.area = r.sizeX * r.sizeY
		r.presents = StringsToIntSlice(rcc[1])
		regions = append(regions, r)
	}

	return presents, regions
}

func StringsToIntSlice(inputText string) []int {
	dataSetStr := strings.Fields(inputText)
	var dataSet []int
	for _, s := range dataSetStr {
		if i, err := strconv.Atoi(s); err == nil {
			dataSet = append(dataSet, i)
		}
	}
	return dataSet
}

func SplitByEmptyNewlineToSlices(str string) [][]string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	strGroups := regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)
	var ret [][]string
	for _, group := range strGroups {
		splitGroup := SplitByLines(group)
		ret = append(ret, splitGroup)
	}
	return ret
}

func SplitByLines(str string) []string {
	return strings.Split(strings.TrimSpace(str), "\n")
}
