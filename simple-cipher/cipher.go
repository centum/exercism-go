package cipher

import "strings"

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type SimpleCipher struct {
	key string
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
			d := rune(c.key[pos%len(c.key)]) - minChar
			if enc {
				r += d
			} else {
				r -= d
			}
			if r > maxChar {
				r += diffRight
			} else if r < minChar {
				r += diffLeft
			}
			res = append(res, r)
			pos++
		}
	}
	return string(res)
}

func (c SimpleCipher) Encode(in string) string {
	return c.convert(in, true)
}

func (c SimpleCipher) Decode(in string) string {
	return c.convert(in, false)
}

func NewCaesar() Cipher {
	return NewVigenere(string('a' + 3))
}

func NewShift(distance int) Cipher {
	if distance >= 26 || distance <= -26 || distance == 0 {
		return nil
	}
	n := 'a' + distance
	if n > maxChar {
		n += diffRight
	} else if n < minChar {
		n += diffLeft
	}
	return NewVigenere(string(n))
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
