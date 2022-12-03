package collfunc_test

import (
	"fmt"
	"testing"

	"github.com/moshenahmias/collfunc"
)

func TestToSlice(t *testing.T) {
	if collfunc.Compare[int](collfunc.ToSlice[int]([]int{0, 1, 2, 3, 4, 5, 6, 7}), []int{0, 1, 2, 3, 4, 5, 6, 7}) != 0 {
		t.FailNow()
	}
}

func TestToMap(t *testing.T) {
	m := collfunc.ToMap([]int{0, 1}, func(k int) string { return fmt.Sprint(k) })

	for k, v := range m {
		if v != fmt.Sprint(k) {
			t.FailNow()
		}
	}
}

func TestCreateMap(t *testing.T) {
	m := collfunc.CreateMap[int, string]([]int{0, 1}, []string{"0", "1"})

	for k, v := range m {
		if v != fmt.Sprint(k) {
			t.FailNow()
		}
	}
}

func TestToChan(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	ch := collfunc.ToChan[int](s)

	for i := range ch {
		if i != s[i] {
			t.FailNow()
		}
	}
}
