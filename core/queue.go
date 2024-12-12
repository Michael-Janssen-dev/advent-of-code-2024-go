package core

type Queue[T any] struct {
	List []T
}

func (q *Queue[T]) Push(item T) {
	q.List = append(q.List, item)
}

func (q *Queue[T]) Pop() T {
	item := q.List[0]
	q.List = q.List[1:]
	return item
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.List) == 0
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{make([]T, 0)}
}