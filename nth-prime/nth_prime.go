package prime

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	res := 1
	k := 0

LOOP0:
	for {
		res++
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
