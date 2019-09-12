// Package letter implement function ConcurrentFrequency
package letter

// ConcurrentFrequency return count the frequency of letters in texts
func ConcurrentFrequency(listText []string) FreqMap {
	ch := make(chan FreqMap, 10)

	for _, s := range listText {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	m := FreqMap{}
	for range listText {
		for k, v := range <-ch {
			m[k] += v
		}
	}

	return m
}
