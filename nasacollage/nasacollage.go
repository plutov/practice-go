package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/shogg/practice-go/nasacollage/apod"
	"github.com/shogg/practice-go/nasacollage/collage"
)

func main() {

	if len(os.Args) == 1 {
		usage()
		return
	}

	switch os.Args[1] {
	case "scrape":
		scrape()
	case "solve":
		solve()
	default:
		usage()
	}
}

func usage() {

	fmt.Println(`

USAGE
	nasacollage scrape
	nasacollage solve <dir> <ground row size>

DESCRIPTION
	1. Scrape image urls and redirect into a text file

		$ nasacollage scrape > urls.txt

	2. Download images

		$ mkdir images; cd images
		$ wget -i ../urls.txt

	3. delete logos and non-image files (for instance *.svf)

	4. Generate collages (execution will take years
		to come to an end, interrupt sometime)

		$ mkdir ../collages; cd ../collages
		$ nasacollage solve ../images 1
		$ nasacollage solve ../images 2
		$ nasacollage solve ../images 3
		$ nasacollage solve ../images 4`)
}

func scrape() {
	err := apod.ScrapeImageURLs(
		"https://apod.nasa.gov/apod/archivepix.html",
		func(s string) { fmt.Println(s) })
	if err != nil {
		log.Fatal(err)
	}
}

func solve() {

	if len(os.Args) != 4 {
		usage()
		return
	}

	dir := os.Args[2]
	resData, err := collage.ListDir(dir)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		usage()
		return
	}

	groundRowSize, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		usage()
		return
	}

	progress := func(current, max int) {
		log.Printf("10^%.4f: 10^%.4f (%.8f %%)\n",
			math.Log10(float64(max)),
			math.Log10(float64(current)),
			float64(current*100)/float64(max))
	}

	collage.NewSolver(resData, progress).Solve(groundRowSize)
}
