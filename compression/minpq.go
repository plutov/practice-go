package compression

import (
	"fmt"
)

type Orderable interface {
	Less(other Orderable) bool
}

type MinPQ[T Orderable] struct {
	arr []T
}

func NewMinPQ[T Orderable]() *MinPQ[T] {
	return &MinPQ[T]{
		arr: make([]T, 0),
	}
}

func (pq *MinPQ[T]) IsEmpty() bool {
	return len(pq.arr) == 0
}

func (pq *MinPQ[T]) Size() int {
	return len(pq.arr)
}

func (pq *MinPQ[T]) Min() (T, error) {
	if len(pq.arr) == 0 {
		var zero T
		return zero, fmt.Errorf("priority queue is empty")
	}
	return pq.arr[0], nil
}

func (pq *MinPQ[T]) swap(a, b int) {
	pq.arr[a], pq.arr[b] = pq.arr[b], pq.arr[a]
}

func (pq *MinPQ[T]) rise(k int) {
	for k > 0 {
		parent := (k - 1) / 2
		if !pq.arr[k].Less(pq.arr[parent]) {
			break
		}
		pq.swap(k, parent)
		k = parent
	}
}

func (pq *MinPQ[T]) sink(k int) {
	n := len(pq.arr)
	for 2*k+1 < n {
		left := 2*k + 1
		right := left + 1
		smallest := left

		if right < n && pq.arr[right].Less(pq.arr[left]) {
			smallest = right
		}

		if !pq.arr[smallest].Less(pq.arr[k]) {
			break
		}

		pq.swap(k, smallest)
		k = smallest
	}
}

func (pq *MinPQ[T]) Insert(x T) {
	pq.arr = append(pq.arr, x)
	pq.rise(len(pq.arr) - 1)
}

func (pq *MinPQ[T]) DeleteMin() (T, error) {
	if len(pq.arr) == 0 {
		var zero T
		return zero, fmt.Errorf("priority queue is empty")
	}
	x := pq.arr[0]
	pq.arr[0] = pq.arr[len(pq.arr)-1]
	pq.arr = pq.arr[:len(pq.arr)-1]

	pq.sink(0)

	return x, nil
}
