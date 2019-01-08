package collage_test

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/shogg/practice-go/nasacollage/collage"
)

func TestSolve(t *testing.T) {

	// skip if travis
	if _, ok := os.LookupEnv("TRAVIS"); ok {
		t.Skip("travis build")
	}

	resData, err := collage.ListDir("../../../tmp/apod")
	if err != nil {
		t.Fatal(err)
	}

	progress := func(current, max int) {
		fmt.Printf("10^%.4f: 10^%.4f (%.8f %%)\n",
			math.Log10(float64(max)),
			math.Log10(float64(current)),
			float64(current*100)/float64(max))
	}

	groundSize, images := collage.NewSolver(resData, progress).Solve(1)
	if groundSize == 0 || len(images) == 0 {
		t.Error("no result")
	}
}
