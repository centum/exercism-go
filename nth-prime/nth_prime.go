package prime

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	if n == 1 {
		return 2, true
	}

	prime := 1

LOOP0:
	for {
		prime += 2
		for i := 3; i < prime; i += 2 {
			if prime%i == 0 {
				continue LOOP0
			}
		}

		if n--; n == 1 {
			return prime, true
		}
	}
}
