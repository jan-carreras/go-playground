package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) (rsp []string) {
	for _, candidate := range candidates {
		if isAnagram(subject, candidate) {
			rsp = append(rsp, candidate)
		}
	}

	return rsp
}

func isAnagram(subject string, candidate string) bool {
	if strings.EqualFold(subject, candidate) {
		return false
	}
	subject, candidate = strings.ToLower(subject), strings.ToLower(candidate)
	return strings.EqualFold(sortStr(subject), sortStr(candidate))
}

func sortStr(subject string) string {
	sub := make([]rune, 0, len(subject))
	for _, c := range subject {
		sub = append(sub, c)
	}

	sort.Slice(sub, func(i, j int) bool {
		return sub[i] < sub[j]
	})

	return string(sub)
}
