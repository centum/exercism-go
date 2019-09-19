package prime

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	if n == 1 {
		return 2, true
	}

	p := 1

LOOP0:
	for {
		p += 2
		for i := 3; i < p; i += 2 {
			if p%i == 0 {
				continue LOOP0
			}
		}

		if n--; n == 1 {
			return p, true
		}
	}
}
