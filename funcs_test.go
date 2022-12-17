package collfunc_test

import (
	"fmt"
	"testing"

	"github.com/moshenahmias/collfunc"
)

func TestCompareEqual(t *testing.T) {
	first := []int{0, 1, 2, 3}
	second := []int{0, 1, 2, 3}

	cmp := collfunc.Compare[int](first, second)
	if cmp != 0 {
		t.Fatal("unexpected result", cmp)
	}
}

func TestCompareFirstGreater(t *testing.T) {
	first := []int{0, 1, 3, 4}
	second := []int{0, 1, 2, 3}

	cmp := collfunc.Compare[int](first, second)
	if cmp >= 0 {
		t.Fatal("unexpected result", cmp)
	}
}

func TestCompareSecondGreater(t *testing.T) {
	first := []int{0, 1, 2, 3}
	second := []int{0, 1, 3, 4}

	cmp := collfunc.Compare[int](first, second)
	if cmp <= 0 {
		t.Fatal("unexpected result", cmp)
	}
}

func TestZip(t *testing.T) {
	first := []int{0, 2, 4, 6}
	second := []int{1, 3, 5, 7}

	if collfunc.Compare[int](collfunc.Zip[int](first, second), []int{0, 1, 2, 3, 4, 5, 6, 7}) != 0 {
		t.FailNow()
	}
}

func TestConcat(t *testing.T) {
	first := []int{0, 1, 2, 3}
	second := []int{4, 5, 6, 7}

	if collfunc.Compare[int](collfunc.Concat[int]([]int{}, first, []int{}, second, []int{}), []int{0, 1, 2, 3, 4, 5, 6, 7}) != 0 {
		t.FailNow()
	}
}

func TestIn(t *testing.T) {
	if !collfunc.In(2, []int{0, 1, 2, 3}) {
		t.FailNow()
	}

	if collfunc.In(6, []int{0, 1, 2, 3}) {
		t.FailNow()
	}
}

func TestFilter(t *testing.T) {
	if collfunc.Compare[int](collfunc.Filter([]int{0, 1, 2, 3, 4, 5, 6, 7}, func(i int) bool { return i%2 == 0 }), []int{0, 2, 4, 6}) != 0 {
		t.FailNow()
	}
}

func TestMap(t *testing.T) {
	if collfunc.Compare[int](collfunc.Map([]int{0, 1, 2, 3, 4, 5, 6, 7}, func(i int) int { return i + 1 }), []int{1, 2, 3, 4, 5, 6, 7, 8}) != 0 {
		t.FailNow()
	}
}

func TestEnumerate(t *testing.T) {
	iter := collfunc.Enumerate[int]([]int{0, 1, 2, 3, 4, 5, 6, 7})

	for i, v, ok := iter(); ok; i, v, ok = iter() {
		if i != v {
			t.FailNow()
		}
	}
}

func TestSink(t *testing.T) {
	if collfunc.Compare[string](
		collfunc.Sink(func(values ...int) string { return fmt.Sprint(values[0] + values[1]) },
			[]int{0, 1, 2},
			[]int{3, 4, 5}),
		[]string{"3", "5", "7"}) != 0 {
		t.FailNow()
	}
}

func TestAccumulate(t *testing.T) {
	if collfunc.Accumulate([]int{0, 1, 2}, func(cur, mem int) int { return cur + mem }) != 3 {
		t.FailNow()
	}
}

func TestOne(t *testing.T) {
	if collfunc.One[int]([]int{3, 2, 1}) != 3 {
		t.FailNow()
	}
}

func TestOneEndOfIteration(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.FailNow()
		}
	}()

	if collfunc.One[int](collfunc.Empty[int]()) != 3 {
		t.FailNow()
	}
}
