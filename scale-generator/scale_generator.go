// Package scale implement function Scale.
package scale

import "strings"

var sharpNotes = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatNotes = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var startFlatNotes = map[string]struct{}{
	"F":  struct{}{},
	"Bb": struct{}{},
	"Eb": struct{}{},
	"Ab": struct{}{},
	"Db": struct{}{},
	"Gb": struct{}{},
	"d":  struct{}{},
	"g":  struct{}{},
	"c":  struct{}{},
	"f":  struct{}{},
	"bb": struct{}{},
	"eb": struct{}{},
}

var periodNotes = map[rune]int{
	'm': 0,
	'M': 1,
	'A': 2,
}

// Scale returns list of chromatic scale.
func Scale(tonic, interval string) []string {
	// Get sharp or flat row
	if tonic != "g" && tonic != "d" {
		tonic = strings.Title(tonic)
	}
	var arr *[]string
	if _, ok := startFlatNotes[tonic]; ok {
		arr = &flatNotes
	} else {
		arr = &sharpNotes
	}

	// Find start position
	tonic = strings.Title(tonic)
	var i int
	for k, v := range *arr {
		if tonic == v {
			i = k
			break
		}
	}

	// initial interval if it's empty
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	// Make row
	res := make([]string, 0, 12)
	for _, c := range interval {
		res = append(res, (*arr)[i])
		p, _ := periodNotes[c]
		i = i + p + 1
		if i >= len(*arr) {
			i = i - len(*arr)
		}
	}

	return res
}
