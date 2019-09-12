// Package letter implement function ConcurrentFrequency
package letter

import "sync"

// ConcurrentFrequency return count the frequency of letters in texts
func ConcurrentFrequency(listText []string) FreqMap {
	m := FreqMap{}
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for _, s := range listText {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			for _, r := range s {
				mu.Lock()
				m[r]++
				mu.Unlock()
			}
		}(s)
	}
	wg.Wait()
	return m
}
