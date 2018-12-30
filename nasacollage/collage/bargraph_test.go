package collage_test

import (
	"reflect"
	"testing"

	"github.com/shogg/practice-go/nasacollage/collage"
)

func TestHighLowIndex(t *testing.T) {

	bars := collage.BarGraph{
		{W: 10, H: 20}, {W: 10, H: 10}, {W: 10, H: 30},
	}

	low := bars.LowIndex()
	if low != 1 {
		t.Errorf("low: want %d got %d", 1, low)
	}

	high := bars.HighIndex()
	if high != 2 {
		t.Errorf("high: want %d got %d", 2, low)
	}
}

func TestStack(t *testing.T) {

	bars := collage.BarGraph{
		{W: 10, H: 20}, {W: 10, H: 10}, {W: 10, H: 30},
	}

	tests := []struct {
		name         string
		stackIndexes []int
		stackBars    []collage.Bar
		expected     collage.BarGraph
	}{
		{"split bar",
			[]int{1},
			[]collage.Bar{{W: 5, H: 30}},
			collage.BarGraph{
				{W: 10, H: 20}, {W: 5, H: 40}, {W: 5, H: 10}, {W: 10, H: 30},
			},
		},
		{"grow bar",
			[]int{1},
			[]collage.Bar{{W: 10, H: 30}},
			collage.BarGraph{
				{W: 10, H: 20}, {W: 10, H: 40}, {W: 10, H: 30},
			},
		},
		{"split bar, merge left",
			[]int{1},
			[]collage.Bar{{W: 5, H: 10}},
			collage.BarGraph{
				{W: 15, H: 20}, {W: 5, H: 10}, {W: 10, H: 30},
			},
		},
		{"grow bar, merge left",
			[]int{1},
			[]collage.Bar{{W: 10, H: 10}},
			collage.BarGraph{
				{W: 20, H: 20}, {W: 10, H: 30},
			},
		},
		{"grow bar, merge right",
			[]int{1},
			[]collage.Bar{{W: 10, H: 20}},
			collage.BarGraph{
				{W: 10, H: 20}, {W: 20, H: 30},
			},
		},
		{"merge left right",
			[]int{0, 1},
			[]collage.Bar{{W: 10, H: 10}, {W: 10, H: 20}},
			collage.BarGraph{
				{W: 30, H: 30},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			barsTmp := make(collage.BarGraph, len(bars))
			copy(barsTmp, bars)

			for i := range test.stackIndexes {
				barsTmp.Stack(test.stackIndexes[i], test.stackBars[i])
			}

			if !reflect.DeepEqual(test.expected, barsTmp) {
				t.Errorf("\nwant\n%v\ngot\n%v\n", test.expected, barsTmp)
			}
		})
	}
}

func TestStackRow(t *testing.T) {

	bars := collage.BarGraph{
		{W: 10, H: 20}, {W: 10, H: 10}, {W: 10, H: 30},
	}

	tests := []struct {
		name       string
		stackIndex int
		stackRow   []collage.Bar
		expected   collage.BarGraph
	}{
		{"split bar",
			1,
			[]collage.Bar{{W: 2, H: 30}, {W: 2, H: 30}},
			collage.BarGraph{
				{W: 10, H: 20}, {W: 4, H: 40}, {W: 6, H: 10}, {W: 10, H: 30},
			},
		},
		{"grow bar",
			1,
			[]collage.Bar{{W: 5, H: 30}, {W: 5, H: 30}},
			collage.BarGraph{
				{W: 10, H: 20}, {W: 10, H: 40}, {W: 10, H: 30},
			},
		},
		{"grow bar, merge left",
			1,
			[]collage.Bar{{W: 5, H: 10}, {W: 5, H: 10}},
			collage.BarGraph{
				{W: 20, H: 20}, {W: 10, H: 30},
			},
		},
		{"split bar, merge right",
			1,
			[]collage.Bar{{W: 5, H: 30}, {W: 5, H: 20}},
			collage.BarGraph{
				{W: 10, H: 20}, {W: 5, H: 40}, {W: 15, H: 30},
			},
		},
		{"split bar, merge left right",
			1,
			[]collage.Bar{{W: 5, H: 10}, {W: 5, H: 20}},
			collage.BarGraph{
				{W: 15, H: 20}, {W: 15, H: 30},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			barsTmp := make(collage.BarGraph, len(bars))
			copy(barsTmp, bars)

			barsTmp.StackRow(test.stackIndex, test.stackRow)

			if !reflect.DeepEqual(test.expected, barsTmp) {
				t.Errorf("\nwant\n%v\ngot\n%v\n", test.expected, barsTmp)
			}
		})
	}
}
