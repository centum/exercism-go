package anagram

import "strings"
import "unicode/utf8"


// Set implement set of rune with counter
type Set struct {
	m map[rune]int
}

// Add added new key or increment key  value
func (s *Set) Add(v rune) {
	if _, ok := s.m[v]; !ok {
		s.m[v] = 0
	}
	s.m[v]++
}

//  Get return value of key
func (s *Set) Get(v rune) int {
	r, _ := s.m[v]
	return r
}

// Equal compare two Set
func (s *Set) Equal(s2 *Set) bool {
	if len(s.m) != len(s2.m) {
		return false
	}

	for k, v := range s2.m {
		if s.Get(k) != v {
			return false
		}
	}
	return true
}

// NewSetFromString return created Set
func NewSetFromString(v string) *Set {
	s := &Set{}
	s.m = make(map[rune]int, utf8.RuneCount([]byte(v)))
	for _, r := range v {
		s.Add(r)
	}

	return s
}

// Detect return list of anagrams
func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)

	subjectSet := NewSetFromString(subject)

	result := make([]string, 0, len(candidates))

	for _, c := range candidates {
		cLower := strings.ToLower(c)
		if cLower == subject {
			continue
		}

		cSet := NewSetFromString(cLower)

		if subjectSet.Equal(cSet) {
			result = append(result, c)
		}
	}

	return result
}
