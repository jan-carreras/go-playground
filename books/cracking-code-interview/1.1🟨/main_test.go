package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

// Is Unique: Implement an algorithm to determine if a string has all unique
// characters. What if you cannot use additional data structures?
func uniqueCharsSupportingDatastructure(s string) bool {
	remember := make(map[rune]bool)
	for _, c := range s {
		if _, seen := remember[c]; seen {
			return false
		}
		remember[c] = true
	}
	return true
}

// Is Unique: Implement an algorithm to determine if a string has all unique
// characters. What if you cannot use additional data structures?
func uniqueCharsBruteforce(s string) bool {
	for i, c := range s {
		for ii, j := range s[i:] {
			if ii == 0 {
				continue // Ignore comparison to itself
			}
			if c == j {
				return false
			}
		}
	}
	return true
}

// Is Unique: Implement an algorithm to determine if a string has all unique
// characters. What if you cannot use additional data structures?
func uniqueCharsBitVector(s string) bool {
	seen := make([]bool, 255)
	for i := 0; i < len(s); i++ {
		char := s[i]
		if found := seen[char]; found {
			return false
		}
		seen[char] = true
	}

	return true
}

type Bitset interface {
	GetBit(i int) bool
	SetBit(i int, value bool)
}

type IntBitset struct {
	btiset [4]uint64
}

func NewIntBitset() *IntBitset {
	return &IntBitset{}
}

func (i2 *IntBitset) GetBit(i int) bool {
	partition := i / 64
	index := i % 64
	return (i2.btiset[partition] & 1 << index) > 0
}

func (i2 *IntBitset) SetBit(i int, value bool) {
	partition := i / 64
	index := i % 64
	if value {
		i2.btiset[partition] |= 1 << index
	} else {
		i2.btiset[partition] &= ^(1 << index)
	}
}

func (i2 *IntBitset) String() string {
	return "bitset: " + strconv.FormatInt(int64(i2.btiset[0]), 2)
}

// Is Unique: Implement an algorithm to determine if a string has all unique
// characters. What if you cannot use additional data structures?
func uniqueCharsSliceOfInt64(s string) bool {
	bs := NewIntBitset()
	for _, char := range s {
		char

	}

	return true
}

func Test_Unique(t *testing.T) {
	require.True(t, uniqueCharsSupportingDatastructure(""))
	require.True(t, uniqueCharsSupportingDatastructure("abc"))
	require.False(t, uniqueCharsSupportingDatastructure("aaa"))
	require.False(t, uniqueCharsSupportingDatastructure("abca"))

	require.True(t, uniqueCharsBruteforce(""))
	require.True(t, uniqueCharsBruteforce("abc"))
	require.False(t, uniqueCharsBruteforce("aaa"))
	require.False(t, uniqueCharsBruteforce("abca"))

	require.True(t, uniqueCharsBitVector(""))
	require.True(t, uniqueCharsBitVector("abc"))
	require.False(t, uniqueCharsBitVector("aaa"))
	require.False(t, uniqueCharsBitVector("abca"))

	bitset := NewIntBitset()
	require.False(t, bitset.GetBit(0))
	bitset.SetBit(0, true)
	require.True(t, bitset.GetBit(0))
	fmt.Println(bitset)
	bitset.SetBit(1, true)
	require.True(t, bitset.GetBit(1))

	fmt.Println(bitset)
	bitset.SetBit(2, true)
	require.True(t, bitset.GetBit(2))
	fmt.Println(bitset)
	bitset.SetBit(7, true)
	fmt.Println(bitset)
	bitset.SetBit(1, false)
	fmt.Println(bitset)
	require.False(t, bitset.GetBit(1))

	// TODO: Write it as a bitmap
}
