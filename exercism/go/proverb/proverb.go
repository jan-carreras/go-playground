// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return nil
	}
	
	rsp := make([]string, 0, len(rhyme)+1)
	for i := 0; i < len(rhyme)-1; i++ {
		rsp = append(rsp, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	rsp = append(rsp, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return rsp
}
