package cipher

import (
	"strings"
)

type shift int32

func NewShift(d int) Cipher {
	if d >= 26 || d <= -26 || d == 0 {
		return nil
	}

	if d < 0 {
		d = 26 + d
	}

	return shift(d)
}

func (s shift) Encode(input string) string {
	input = toASCIILower(input)

	b := strings.Builder{}
	diff := int32(s)
	for _, c := range input {
		b.WriteRune(((c - 'a' + diff) % 26) + 'a')
	}

	return b.String()
}

func (s shift) Decode(input string) string {
	b := strings.Builder{}

	diff := int32(26 - s)
	for _, c := range input {
		b.WriteRune(((c - 'a' + diff) % 26) + 'a')
	}

	return b.String()
}

func toASCIILower(input string) string {
	b := strings.Builder{}
	for _, c := range input {
		if c >= 'a' && c <= 'z' {
			b.WriteRune(c)
		} else if c >= 'A' && c <= 'Z' {
			b.WriteRune(c - 'A' + 'a')
		}
	}

	return b.String()
}

func NewCaesar() Cipher {
	return NewShift(3)
}

type vigenere []int32

func NewVigenere(key string) Cipher {
	valid := false
	for _, c := range key {
		if c == 'a' {
			continue
		} else if c > 'a' && c <= 'z' {
			valid = true
		} else {
			return nil
		}
	}

	if !valid {
		return nil
	}

	b := strings.Builder{}
	for _, c := range toASCIILower(key) {
		b.WriteRune(c - 'a')
	}

	return vigenere(b.String())
}

func (v vigenere) Encode(input string) string {
	b := strings.Builder{}
	for i, c := range toASCIILower(input) {
		s := ((c - 'a' + v[i%len(v)]) % 26) + 'a'
		b.WriteRune(s)
	}

	return b.String()
}

func (v vigenere) Decode(input string) string {
	b := strings.Builder{}
	for i, c := range input {
		diff := (26 - v[i%len(v)]) % 26
		s := ((c - 'a' + diff) % 26) + 'a'
		b.WriteRune(s)
	}

	return b.String()
}
