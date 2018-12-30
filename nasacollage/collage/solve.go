package collage

import (
	"fmt"
	"log"
	"sort"
	"time"
)

// Solver builds rectangular collages
type Solver struct {
	resData          []Imgres
	progress         *Progress
	progressCallback func(int, int)

	solutionIndexes    []int
	solutionArea       int
	solutionGroundSize int
}

// NewSolver creates a solver
func NewSolver(
	resData []Imgres,
	progressCallback func(int, int)) *Solver {

	sortByWidth := func(resData []Imgres) []Imgres {
		sort.SliceStable(resData, func(i, j int) bool {
			return resData[i].W < resData[j].W
		})
		return resData
	}

	return &Solver{
		resData:          sortByWidth(resData),
		solutionArea:     10e5,
		progressCallback: progressCallback,
	}
}

// Solve build rectangular collages and save to disk.
func (s *Solver) Solve(groundRowSize int) (int, []Imgres) {

	// iterate variations of k ground row images out of n images overall.
	// example: 8000!/(8000-3)! ~= 10^11.7 variations
	n := len(s.resData)
	k := groundRowSize
	s.progress = NewProgress(NumVariations(n, k), s.progressCallback)
	Variations(n, k, s.solve)

	// build solution data from indexes
	solutionImages := make([]Imgres, len(s.solutionIndexes))
	for i, ii := range s.solutionIndexes {
		solutionImages[i] = s.resData[ii]
	}

	return s.solutionGroundSize, solutionImages
}

func (s *Solver) solve(groundRow []int) {

	s.progress.Inc()

	// use a bar graph to keep track of stacked heights
	bars := NewBarGraph(1)
	s.solveRecursively(bars, 0, nil, groundRow, len(groundRow))
}

func (s *Solver) solveRecursively(
	bars BarGraph, gapIndex int,
	imgIndexes, newImgIndexes []int,
	groundSize int) {

	if len(imgIndexes) > 11 {
		return
	}

	// abort if an image got used twice
	if !Disjoint(imgIndexes, newImgIndexes) {
		return
	}

	// copy images and append new images
	imgIndexesCopy := make([]int, len(imgIndexes), len(imgIndexes)+len(newImgIndexes))
	copy(imgIndexesCopy, imgIndexes)
	imgIndexes = imgIndexesCopy
	imgIndexes = append(imgIndexes, newImgIndexes...)

	// place images in the lowest bar
	bars = s.placeImagesInGap(bars, gapIndex, newImgIndexes)

	// rectangle found
	if len(bars) == 1 {
		area := bars[0].W * bars[0].H

		if /*area < s.solutionArea &&*/ len(imgIndexes) > 10 &&
			(s.solutionIndexes == nil || s.solutionIndexes[0] != imgIndexes[0]) {

			s.solutionArea = area
			s.solutionIndexes = imgIndexes
			s.solutionGroundSize = groundSize
			log.Println(groundSize, imgIndexes, bars[0].W, "x", bars[0].H)
			s.writePNG()
			return
		}
	}

	gapIndex = bars.LowIndex()
	s.variationsOfFittingImages(
		bars, gapIndex,
		imgIndexes, groundSize,
		s.solveRecursively)
}

func (s *Solver) variationsOfFittingImages(
	bars BarGraph, gapIndex int,
	imgIndexes []int, groundSize int,
	f func(BarGraph, int, []int, []int, int)) {

	s.iterateFittingImages(bars[gapIndex].W, func(newImgIndexes []int) {
		Permutations(newImgIndexes, func(newImgIndexes []int) {
			f(bars, gapIndex, imgIndexes, newImgIndexes, groundSize)
		})
	})
}

func (s *Solver) iterateFittingImages(width int, f func([]int)) {
	newImgIndexes := make([]int, 0, 5)
	s.iterateFittingImagesRecurively(width, f, newImgIndexes)
}

func (s *Solver) iterateFittingImagesRecurively(width int, f func([]int), newImgIndexes []int) {

	for i := 0; i < s.indexWider(width); i++ {

		if !Disjoint(newImgIndexes, []int{i}) {
			continue
		}

		newImgIndexes = append(newImgIndexes, i)
		gap := width - s.resData[i].W

		if gap > 0 {
			s.iterateFittingImagesRecurively(gap, f, newImgIndexes)
		} else {
			f(newImgIndexes)
		}

		newImgIndexes = newImgIndexes[:len(newImgIndexes)-1]
	}
}

func (s *Solver) indexWider(width int) int {
	return sort.Search(len(s.resData), func(i int) bool {
		return s.resData[i].W > width
	})
}

func (s *Solver) placeImagesInGap(bars BarGraph, gapIndex int, newImgIndexes []int) BarGraph {

	// copy bars before changing them
	barsTmp := make(BarGraph, len(bars))
	copy(barsTmp, bars)
	bars = barsTmp

	row := make([]Bar, len(newImgIndexes))
	for i, ii := range newImgIndexes {
		row[i] = Bar{W: s.resData[ii].W, H: s.resData[ii].H}
	}

	bars.StackRow(gapIndex, row)
	return bars
}

func (s *Solver) writePNG() {

	solutionImages := make([]Imgres, len(s.solutionIndexes))
	for i, ii := range s.solutionIndexes {
		solutionImages[i] = s.resData[ii]
	}

	filename := fmt.Sprintf("collage_%d_%d.png", time.Now().Unix(), s.solutionArea)
	if err := WriteCollagePNG(filename, s.solutionGroundSize, solutionImages); err != nil {
		panic(err)
	}
}
