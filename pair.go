package fun

type Pair[T any, S any] struct {
	First  T
	Second S
}

func (p Pair[T, S]) Unpack() (T, S) {
	return p.First, p.Second
}
