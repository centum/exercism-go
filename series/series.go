package series

func All(n int, s string) []string {
	lenStr := len(s)

	res := make([]string, 0, lenStr-n+1)

	for i := 0; i+n <= lenStr; i++ {
		res = append(res, s[i:i+n])
	}

	return res
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

func First(n int, s string) (string, bool) {
	r := All(n, s)
	if r == nil || len(r) == 0 {
		return "", false
	}
	return r[0], true
}
