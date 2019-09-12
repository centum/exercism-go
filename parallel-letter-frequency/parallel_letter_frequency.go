// Package letter implement function ConcurrentFrequency
package letter

// ConcurrentFrequency return count the frequency of letters in texts
func ConcurrentFrequency(listText []string) FreqMap {
	ch := make(chan FreqMap, 10)

	for _, s := range listText {
		go func(s string) {
			m := FreqMap{}
			for _, r := range s {
				m[r]++
			}
			ch <- m
		}(s)
	}

	m := FreqMap{}
	for i := len(listText); i > 0; i-- {
		for k, v := range <-ch {
			m[k] += v
		}
	}

	return m
}
