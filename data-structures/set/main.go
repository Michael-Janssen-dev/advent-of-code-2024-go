package set

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Copy() Set[T] {
	newSet := NewSet[T]()
	for k := range s {
		newSet.Add(k)
	}
	return newSet
}
