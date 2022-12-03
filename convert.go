package collfunc

// ToSlice initializes and returns a new slice
// from the given collection
func ToSlice[T any](collection any) []T {
	var slice []T

	iter := Iterate[T](collection)

	for v, ok := iter(); ok; v, ok = iter() {
		slice = append(slice, v)
	}

	return slice
}

// ToMap initializes and returns a new map
// from the given collection, a mapping function
// is called to map a value for every key
func ToMap[K comparable, V any](collection any, f func(K) V) map[K]V {
	m := make(map[K]V)

	iter := Iterate[K](collection)

	for k, ok := iter(); ok; k, ok = iter() {
		m[k] = f(k)
	}

	return m
}

// CreateMap initializes and returns a new map
// with keys from the keys collection and corresponding
// values from the values collection
func CreateMap[K comparable, V any](keys any, values any) map[K]V {
	iter := Iterate[V](values)

	return ToMap(keys, func(K) V {
		v, _ := iter()
		return v
	})
}

// ToChan writes all given collection items to a newly
// created channel
func ToChan[T any](collection any) <-chan T {
	ch := make(chan T)

	iter := Iterate[T](collection)

	go func() {
		defer close(ch)
		for v, ok := iter(); ok; v, ok = iter() {
			ch <- v
		}
	}()

	return ch
}
