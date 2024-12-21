// Adapted from https://pkg.go.dev/container/heap#example-package-PriorityQueue
package core

import (
	"container/heap"
)

type PriorityQueueItem interface {
	Priority() int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T PriorityQueueItem] []T

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority() > pq[j].Priority()
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(T)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (h *PriorityQueue[T]) PushSafe(x T) {
	heap.Push(h, x)
}

func (h *PriorityQueue[T]) PopSafe() T {
	return heap.Pop(h).(T)
}

func NewPriorityQueue[T PriorityQueueItem]() PriorityQueue[T] {
	return make(PriorityQueue[T], 0)
}
