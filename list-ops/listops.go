// Package listops implement basic function  on list int.
package listops

type binFunc func(int, int) int

type IntList []int

// Foldl returns reduce values of IntList.
func (l IntList) Foldl(fn binFunc, initial int) int {
	for _, v := range l {
		initial = fn(initial, v)
	}
	return initial
}

// Foldr returns reduce values of reverse IntList.
func (l IntList) Foldr(fn binFunc, initial int) int {
	for i := len(l) - 1; i>=0; i-- {
		initial = fn(l[i], initial)
	}
	return initial
}

type predFunc func(n int) bool

// Filter returns filtering IntList.
func (l IntList) Filter(fn predFunc) IntList {
	newList := make(IntList, 0)
	for _, v := range l {
		if fn(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// Length returns length IntList.
func (l IntList) Length() (c int) {
	for range l {
		c++
	}
	return c
}

type unaryFunc func(x int) int

// Map returns mapping IntList.
func (l IntList) Map(fn unaryFunc) IntList {
	res := make(IntList, 0, l.Length())
	for _, v := range l {
		res = append(res, fn(v))
	}
	return res
}

// Reverse returns reverse IntList.
func (l IntList) Reverse() IntList {
	newList := make(IntList, 0, l.Length())
	for i := len(l) - 1; i>=0; i-- {
		newList = append(newList, l[i])
	}

	return newList
}

// Append returns IntList after append IntList.
func (l IntList) Append(appendList IntList) IntList {
	res := make(IntList, 0, l.Length() + appendList.Length())
	res = append(res, l...)
	res = append(res, appendList...)
	return res
}

// Concat returns list after append list of IntList.
func (l IntList) Concat(concatList []IntList) IntList {
	res := make(IntList, 0, l.Length())
	res = append(res, l...)
	for _, ll := range concatList {
		res = append(res, ll...)
	}
	return res
}

