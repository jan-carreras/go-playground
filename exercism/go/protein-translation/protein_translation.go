package protein

import (
	"errors"
)

var ErrStop = errors.New("stop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

func FromRNA(rna string) ([]string, error) {
	codons := make([]string, 0, len(rna)/3)

	for i := 0; i < len(rna); i += 3 {
		codon, err := FromCodon(rna[i : i+3])
		if errors.Is(err, ErrStop) {
			break
		} else if err != nil {
			return nil, err
		}

		codons = append(codons, codon)
	}

	return codons, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}

	return "", ErrInvalidBase
}
