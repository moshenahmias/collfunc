package collfunc

import "golang.org/x/exp/constraints"

// In returns true if the given item is found within the collection,
// it returns false otherwise.
func In[T comparable](item T, collection any) bool {
	iter := Iterate[T](collection)

	for v, ok := iter(); ok; v, ok = iter() {
		if v == item {
			return true
		}
	}

	return false
}

// Filter filters a collection using a predicate, it will return
// the result as an iterator.
func Filter[T any](collection any, f func(T) bool) func() (T, bool) {
	iter := Iterate[T](collection)

	return func() (T, bool) {
		for {
			v, ok := iter()
			if !ok {
				return v, ok
			}

			if f(v) {
				return v, ok
			}
		}
	}
}

// Map every item in the collection using the mapping function,
// returns an iterator as the result.
func Map[T, S any](collection any, f func(T) S) func() (S, bool) {
	iter := Iterate[T](collection)

	return func() (S, bool) {
		v, ok := iter()

		if !ok {
			return *new(S), ok
		}

		return f(v), ok
	}
}

// Enumerate returns an indexed iterator for the given collection.
func Enumerate[T any](collection any) func() (int, T, bool) {
	iter := Iterate[T](collection)
	i := 0
	return func() (int, T, bool) {
		v, ok := iter()
		i++
		return i - 1, v, ok
	}
}

// Sink maps values from multiple collections into a single collection,
// returned as an iterator.
func Sink[T, S any](f func(values ...T) S, collections ...any) func() (S, bool) {
	var iters []func() (T, bool)

	for _, collection := range collections {
		iters = append(iters, Iterate[T](collection))
	}

	done := true

	return func() (S, bool) {
		var values []T

		for _, iter := range iters {
			v, ok := iter()

			if ok {
				done = false
			}

			values = append(values, v)
		}

		if done {
			return *new(S), false
		}

		done = true

		return f(values...), true
	}
}

// Accumulate calculation on collection items.
func Accumulate[T any](collection any, f func(cur, mem T) T) T {
	var mem T

	iter := Iterate[T](collection)

	for v, ok := iter(); ok; v, ok = iter() {
		mem = f(v, mem)
	}

	return mem
}

// Zip returns an iterator that iterates over the two collections
// in a zipped order.
func Zip[T any](first, second any) func() (T, bool) {
	iter := [2]func() (T, bool){Iterate[T](first), Iterate[T](second)}
	i := 0
	vi, oki := iter[i]()

	return func() (T, bool) {
		i = (i + 1) % 2
		vj, okj := iter[i]()
		v := vi

		var ok bool

		if oki {
			ok = oki
		} else {
			ok = okj
		}

		vi, oki = vj, okj

		return v, ok
	}
}

// Concat returns an iterator that iterates over all the
// collections.
func Concat[T any](collections ...any) func() (T, bool) {
	if len(collections) == 0 {
		return Empty[T]()
	}

	i := 0
	iter := Iterate[T](collections[i])

	return func() (T, bool) {

		for {

			v, ok := iter()

			if !ok && i+1 < len(collections) {
				i++
				iter = Iterate[T](collections[i])
				continue
			}

			return v, ok
		}
	}
}

// Compare two collections, negative if first is bigger, positive
// if second, zero for equality.
func Compare[T constraints.Ordered](first, second any) int {
	firstIter := Iterate[T](first)
	secondIter := Iterate[T](second)

	vi, oki := firstIter()
	vj, okj := secondIter()

	for oki || okj {
		if oki && !okj {
			return -1
		}

		if !oki && okj {
			return 1
		}

		if !oki && !okj {
			return 0
		}

		if vi > vj {
			return -1
		}

		if vi < vj {
			return 1
		}

		vi, oki = firstIter()
		vj, okj = secondIter()
	}

	return 0
}

func One[T any](collection any) T {
	iter := Iterate[T](collection)

	if v, ok := iter(); ok {
		return v
	}

	panic("collfunc: end of iteration")
}
