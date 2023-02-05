package strand

import (
	"fmt"
	"strings"
)

func ToRNA(dna string) string {
	b := strings.Builder{}
	b.Grow(len(dna))

	for _, c := range dna {
		switch c {
		case 'G':
			b.WriteRune('C')
		case 'C':
			b.WriteRune('G')
		case 'T':
			b.WriteRune('A')
		case 'A':
			b.WriteRune('U')
		default:
			panic(fmt.Errorf("unknown nucleotides %q", c))
		}
	}

	return b.String()
}
