package prime

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	if n == 1 {
		return 2, true
	}

	res := 1
	k := 1

LOOP0:
	for {
		res += 2
		for i := res - 1; i >= 2; i-- {
			if res%i == 0 {
				continue LOOP0
			}
		}
		k++
		if k == n {
			return res, true
		}
	}
}
