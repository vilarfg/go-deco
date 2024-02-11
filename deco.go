package deco

func Chain[T any](decorators ...Decorator[T]) Decorator[T] { return extend(nil, decorators) }

type Decorator[T any] func(T) T

func (m Decorator[T]) Extend(decorators ...Decorator[T]) Decorator[T] { return extend(m, decorators) }

func (m Decorator[T]) Apply(t T) T { return m(t) }

func extend[T any](original Decorator[T], ds []Decorator[T]) Decorator[T] {
	var (
		i = -1
		d Decorator[T]
	)
	for _, d = range ds {
		if d != nil {
			i++
			ds[i] = d
		}
	}
	switch i {
	case -1:
		if original == nil {
			return identity
		}
		return func(t T) T { return original(t) }
	case 0:
		d = ds[0]
		if original == nil {
			return func(t T) T { return d(t) }
		}
		return func(t T) T { return original(d(t)) }
	}
	if original == nil {
		return func(t T) T {
			var j = i
			for ; j >= 0; j-- {
				t = ds[j](t)
			}
			return t
		}
	}
	return func(t T) T {
		var j = i
		for ; j >= 0; j-- {
			t = ds[j](t)
		}
		return original(t)
	}
}

func identity[T any](t T) T { return t }
