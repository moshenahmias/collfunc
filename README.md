# collfunc

`Iterate`, `Compare`, `Filter`, `Map`, etc... collection (arrays, slices, maps and custom types) functions!

``` go
func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	iter := collfunc.Iterate[int](s)
	for v, ok := iter(); ok; v, ok = iter() {
	    fmt.Println(v)
	}
}
```
Output:
```
0
1
2
3
4
5
```

`Zip`:

``` go
func main() {
	first := []int{0, 2, 4, 6}
	second := []int{1, 3, 5, 7}
	fmt.Println(collfunc.ToSlice[int](collfunc.Zip[int](first, second)))
}
```
Output:
```
[0, 1, 2, 3, 4, 5, 6, 7]
```


Support for custom collections:

``` go
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

func main() {
	var ci customIterable = 5
	if collfunc.Compare[int](ci, []int{0, 1, 2, 3, 4}) != 0 {
		panic(":(")
	}
}
```
