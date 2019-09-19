package cipher

import (
	"strings"
)

type SimpleCipher struct {
	distance int
	key      string
}

const (
	minChar   = 97
	maxChar   = 122
	diffRight = minChar - maxChar - 1
	diffLeft  = maxChar - minChar + 1
)

func (c SimpleCipher) convert(in string, enc bool) string {
	in = strings.ToLower(in)
	res := make([]rune, 0, len(in))
	pos := 0
	for _, r := range in {
		if r >= minChar && r <= maxChar {
			n := r + c.getDistance(enc, pos)
			if n > maxChar {
				n += diffRight
			} else if n < minChar {
				n += diffLeft
			}
			res = append(res, n)
			pos++
		}
	}
	return string(res)
}

func (c SimpleCipher) getDistance(enc bool, pos int) int32 {
	if c.distance != 0 {
		if enc {
			return int32(c.distance)
		}
		return -int32(c.distance)
	}
	pos = pos % len(c.key)
	d := rune(c.key[pos]) - minChar
	if enc {
		return d
	}
	return -d
}

func (c SimpleCipher) Encode(in string) string {
	return c.convert(in, true)
}

func (c SimpleCipher) Decode(in string) string {
	return c.convert(in, false)
}

func NewCaesar() Cipher {
	return SimpleCipher{distance: 3}
}

func NewShift(distance int) Cipher {
	if distance >= 26 || distance <= -26 || distance == 0 {
		return nil
	}
	return SimpleCipher{distance: distance}
}

func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}
	onlyA := true
	for _, r := range key {
		if r < minChar || r > maxChar {
			return nil
		}
		if r != 'a' {
			onlyA = false
		}
	}
	if onlyA {
		return nil
	}
	return SimpleCipher{key: key}
}
