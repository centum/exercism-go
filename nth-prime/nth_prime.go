package prime

import "fmt"

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	res := 0

LOOP0:
	for i := 2; res < n; i++ {
		for k := i; k > 1; k-- {
			if i%k == 0 {
				continue LOOP0
			}
		}
		res++
		fmt.Printf("%d -> %d\n", res, i)
	}

	return res, true
}
