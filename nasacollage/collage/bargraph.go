package collage

// Bar of a graph
type Bar struct {
	H, W int
}

// BarGraph bar graph
type BarGraph []Bar

// NewBarGraph create new
func NewBarGraph(size int) BarGraph {
	return make(BarGraph, size)
}

// LowIndex index of the lowest bar
func (bars BarGraph) LowIndex() int {
	return Min(len(bars), func(i int) int { return bars[i].H })
}

// HighIndex index of the highest bar
func (bars BarGraph) HighIndex() int {
	return Max(len(bars), func(i int) int { return bars[i].H })
}

// Stack a bar outermost left at index
func (bars *BarGraph) Stack(index int, bar Bar) {
	bars.StackRow(index, []Bar{bar})
}

// StackRow stack a row of bars side by side at index.
// Merge adjacent bars of same height.
func (bars *BarGraph) StackRow(index int, barRow []Bar) {

	for _, b := range barRow {

		// same width
		if (*bars)[index].W == b.W {
			(*bars)[index].H += b.H

			// merge left
			if index > 0 {
				if (*bars)[index-1].H == (*bars)[index].H {
					(*bars)[index-1].W += (*bars)[index].W
					*bars = append((*bars)[:index], (*bars)[index+1:]...)
					index--
				}
			}

			// merge right
			if index < len(*bars)-1 {
				if (*bars)[index].H == (*bars)[index+1].H {
					(*bars)[index].W += (*bars)[index+1].W
					*bars = append((*bars)[:index+1], (*bars)[index+2:]...)
				}
			}
			return
		}

		b.H += (*bars)[index].H
		(*bars)[index].W -= b.W

		// merge left
		if index > 0 {
			if (*bars)[index-1].H == b.H {
				(*bars)[index-1].W += b.W
				continue
			}
		}

		// insert new bar
		*bars = append(*bars, Bar{})
		copy((*bars)[index+1:], (*bars)[index:])
		(*bars)[index] = b
		index++
	}

	if (*bars)[index].W <= 0 {
		*bars = append((*bars)[:index], (*bars)[index+1:]...)
	}
}
