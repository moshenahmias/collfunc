package collfunc_test

import (
	"testing"

	"github.com/moshenahmias/collfunc"
)

func TestEmpty(t *testing.T) {
	if _, ok := collfunc.Empty[any]()(); ok {
		t.FailNow()
	}
}

func TestIterate(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	iter := collfunc.Iterate[int](s)
	for v, ok := iter(); ok; v, ok = iter() {
		if v != s[v] {
			t.FailNow()
		}
	}
}

type customIterable int

func (ci customIterable) Iterator() func() (int, bool) {
	i := 0

	return func() (int, bool) {
		if i < int(ci) {
			i++
			return i - 1, true
		}

		return 0, false
	}
}

func TestIterateCustom(t *testing.T) {
	var ci customIterable = 5
	if collfunc.Compare[int](ci, []int{0, 1, 2, 3, 4}) != 0 {
		t.FailNow()
	}
}
