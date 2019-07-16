package strain

// Ints list of int
type Ints []int

// Lists of slice int
type Lists [][]int

// Strings list of string
type Strings []string

type fnInt func(int) bool
type fnSliceInt func([]int) bool
type fnString func(string) bool

// Keep teturn Ints
func (i Ints) Keep(fn fnInt) Ints {
	res := make(Ints, 0, len(i))
	for _, v := range i {
		if fn(v) {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return Ints(nil)
	}
	return res
}

// Discard return Ints
func (i Ints) Discard(fn fnInt) Ints {
	res := make(Ints, 0, len(i))
	for _, v := range i {
		if !fn(v) {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return Ints(nil)
	}
	return res
}

// Keep return Lists
func (l Lists) Keep(fn fnSliceInt) Lists {
	res := make(Lists, 0, len(l))
	for _, v := range l {
		if fn(v) {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return Lists(nil)
	}
	return res
}

// Keep return Strings
func (s Strings) Keep(fn fnString) Strings {
	res := make(Strings, 0, len(s))
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return Strings(nil)
	}
	return res
}
