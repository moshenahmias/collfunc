package collfunc

import (
	"fmt"
	"reflect"
)

// Iterable types that implement this interface
// can be used with all the funcs in this package.
type Iterable[T any] interface {
	// Iterator returns an interator in the form
	// of a function that returns the next value
	// and a second boolean that signals the end
	// of the iteration.
	Iterator() func() (T, bool)
}

// Iterate returns an iterator for iterating over
// the collection, which can be of type Slice, Array,
// Map, Iterable or iterator (func() (T, bool), will
// be returned as-is).
func Iterate[T any](collection any) func() (T, bool) {
	if iter, ok := collection.(func() (T, bool)); ok {
		return iter
	}

	if iterable, ok := collection.(Iterable[T]); ok {
		return iterable.Iterator()
	}

	var iter func() (reflect.Value, bool)
	value := reflect.ValueOf(collection)

	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		iter = iterateSlice(value)
	case reflect.Map:
		iter = iterateMap(value)
	default:
		panic(fmt.Sprint("collfunc: invalid collection", value.Kind()))
	}

	return func() (T, bool) {
		v, ok := iter()
		return v.Interface().(T), ok
	}
}

// Empty returns an iterator that always returns
// the zero value of T and false.
func Empty[T any]() func() (T, bool) {
	var v T
	return func() (T, bool) {
		return v, false
	}
}

func iterateSlice(value reflect.Value) func() (reflect.Value, bool) {
	i := 0
	length := value.Len()

	return func() (reflect.Value, bool) {
		if i < length {
			v := value.Index(i)
			i++
			return v, true
		} else {
			return reflect.Zero(value.Type().Elem()), false
		}
	}
}

func iterateMap(value reflect.Value) func() (reflect.Value, bool) {
	iter := value.MapRange()
	next := iter.Next()

	return func() (reflect.Value, bool) {
		if next {
			v := iter.Key()
			next = iter.Next()
			return v, true
		} else {
			return reflect.Zero(value.Type().Key()), false
		}
	}
}
