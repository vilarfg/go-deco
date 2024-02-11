// package deco facilitates the chaining of decorators.
package deco

// Chain combines multiple decorators into a single one.
//
// The resulting decorator will apply the provided decorators in reverse order.
func Chain[T any](decorators ...Decorator[T]) Decorator[T] { return extend(nil, decorators) }

// Decorator is a function that takes a value of type `T`
// and returns a potentially modified value of the same type.
type Decorator[T any] func(T) T

// Extends an existing decorator with additional decorators.
//
// Note: decorators are applied in reverse order
func (m Decorator[T]) Extend(decorators ...Decorator[T]) Decorator[T] { return extend(m, decorators) }

// Applies the decorator to a value.
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
	return func(t T) T {
		var j = i
		for ; j >= 0; j-- {
			t = ds[j](t)
		}
		if original == nil {
			return t
		}
		return original(t)
	}
}
