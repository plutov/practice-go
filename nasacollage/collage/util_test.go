package collage_test

import (
	"strconv"
	"testing"

	"github.com/shogg/practice-go/nasacollage/collage"
)

func TestDisjoint(t *testing.T) {

	tests := []struct {
		a, b     []int
		disjoint bool
	}{
		{nil, nil, true},
		{[]int{}, nil, true},
		{[]int{}, []int{}, true},
		{[]int{0}, []int{}, true},
		{[]int{0}, []int{1}, true},
		{[]int{0, 1, 2}, []int{3}, true},
		{[]int{1, 2}, []int{3, 0}, true},
		{[]int{1, 2, 3}, []int{2, 0}, false},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if collage.Disjoint(test.a, test.b) != test.disjoint {
				t.Errorf("%v %v disjoint=%v expected", test.a, test.b, test.disjoint)
			}
		})
	}
}

func TestMinMax(t *testing.T) {

	data := []struct{ x, y int }{{5, 6}, {4, 7}, {3, 8}, {2, 9}}
	getX := func(i int) int { return data[i].x }
	getY := func(i int) int { return data[i].y }

	minX := collage.Min(len(data), getX)
	if minX != 3 {
		t.Errorf("min x = 3 expected, was %d", minX)
	}
	minY := collage.Min(len(data), getY)
	if minY != 0 {
		t.Errorf("min y = 0 expected, was %d", minY)
	}
	maxX := collage.Max(len(data), getX)
	if maxX != 0 {
		t.Errorf("max x = 0 expected, was %d", maxX)
	}
	maxY := collage.Max(len(data), getY)
	if maxY != 3 {
		t.Errorf("max y = 3 expected, was %d", maxY)
	}
}
