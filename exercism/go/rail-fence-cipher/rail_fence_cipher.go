package railfence

import (
	"strings"
)

func Encode(message string, rails int) string {
	message = cleanInput(message)

	seq := sequence(rails)
	groups := make([][]byte, rails)
	for i := 0; i < len(message); i++ {
		groupID := seq[i%len(seq)]
		groups[groupID] = append(groups[groupID], message[i])
	}

	b := strings.Builder{}
	b.Grow(len(message))
	for _, g := range groups {
		b.WriteString(string(g))
	}

	return b.String()
}

func cleanInput(input string) string {
	b := strings.Builder{}
	for _, c := range input {
		if c >= 'A' && c <= 'Z' {
			b.WriteRune(c)
		}
	}

	return b.String()
}

func sequence(rails int) []int {
	sequence := make([]int, 0, rails)
	for i := 0; i < rails; i++ {
		sequence = append(sequence, i)
	}
	for i := rails - 2; i > 0; i-- {
		sequence = append(sequence, i)
	}
	return sequence
}

func Decode(message string, rails int) string {
	// Input WECRLTEERDSOEEFEAOCAIVDEN

	// Challenges: I don't know "how big" is each group

	// Option 1:
	//  - Encode the cyphered input to know how big is each group
	//  - Replace each group with the corresponding part of the string
	//  - Iterate thru the groups using the `sequence` to append character to the buffer
	//  - Remove elements from the list as I go along

	// Option 2: Discover what the groups should be????
	//  - I don't see how I could do that, AND then I would need to operate like Option 1

	// Chosen option 1

	seq := sequence(rails)
	groups := make([][]byte, rails)
	for i := 0; i < len(message); i++ {
		groupID := seq[i%len(seq)]
		groups[groupID] = append(groups[groupID], message[i])
	}

	i := 0
	for gID, g := range groups {
		groups[gID] = []byte(message[i : i+len(g)])
		i += len(g)
	}

	b := strings.Builder{}
	for i := 0; i < len(message); i++ {
		gID := seq[i%len(seq)]
		b.WriteByte(groups[gID][0])
		groups[gID] = groups[gID][1:]
	}

	return b.String()
}
