package queuev1

type Queue[T comparable] struct {
	container []T
}

func QueueOf[T comparable](size int) Queue[T] {
	return Queue[T]{
		container: make([]T, size)[:],
	}
}

func (x *Queue[T]) PushBack(v T) int {
	x.container = append(x.container, v)
	return x.Len() - 1
}

func (x *Queue[T]) Get(idx int) T {
	return x.container[idx]
}

func (x *Queue[T]) Len() int {
	return len(x.container)
}

func (x *Queue[T]) FastRemove(idx int) {
	last := x.Len() - 1

	x.Set(idx, x.Get(last))

	x.container = x.container[:last]
}

func (x *Queue[T]) Set(idx int, v T) {
	x.container[idx] = v
}

func (x *Queue[T]) Contains(fn func(T) bool) bool {
	// for _, t := range x.container {
	// 	if fn(t) {
	// 		return true
	// 	}
	// }

	return false
}

func (x *Queue[T]) Peak() T {
	return x.Get(0)
}

func (x *Queue[T]) Pop() T {
	v := x.Get(0)
	if x.Len() > 1 {
		x.container = x.container[1:]
	} else {
		x.container = []T{}
	}

	return v
}

func (x *Queue[T]) HasNext() bool {
	return x.Len() > 0
}
