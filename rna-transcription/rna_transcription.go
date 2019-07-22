// Package strand implement function ToRNA
package strand

import "strings"

// Mapping DNA to RNA
var mapConvert = map[int32]byte {
	'G': byte('C'),
	'C': byte('G'),
	'T': byte('A'),
	'A': byte('U'),
}

// ToRNA returns rna from dna
func ToRNA(dna string) string {
	rnaBuilder := strings.Builder{}

	for _, c := range dna {
		rnaBuilder.WriteByte(mapConvert[c])
	}

	return rnaBuilder.String()
}
