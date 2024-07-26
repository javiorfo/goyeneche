package optional

type Optional[T any] struct {
	value *T
}

func (o Optional[T]) Get() T {
	if o.value != nil {
		return *o.value
	}
	panic("value does not exist")
}

func (o Optional[T]) OrElseGet(fn func() T) T {
	if o.value != nil {
		return *o.value
	}
	return fn()
}

func (o Optional[T]) OrElse(t T) T {
	if o.value != nil {
		return *o.value
	}
	return t
}

func (o Optional[T]) IsPresent() bool {
	if o.value != nil {
		return true
	}
	return false
}

func (o Optional[T]) IfPresent(fn func(T)) {
	if o.value != nil {
		fn(*o.value)
	}
}

func Of[T any](t T) Optional[T] {
	return Optional[T]{&t}
}

func OfNullable[T any](t *T) Optional[T] {
	return Optional[T]{t}
}
