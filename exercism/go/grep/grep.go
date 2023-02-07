package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	flagPrintLineNumber   = "-n"
	flagPrintFilenameOnly = "-l"
	flagIgnoreCase        = "-i"
	flagExactMatch        = "-x"
	flagInvertedMatch     = "-v"
)

type match struct {
	filename string
	line     int
	text     string
}

func Search(pattern string, flags, files []string) []string {
	return NewGrep(pattern, flags, files).Grep()
}

type Grep struct {
	pattern string
	flags   map[string]bool
	files   []string
}

func NewGrep(pattern string, flags []string, files []string) *Grep {
	mapFlag := make(map[string]bool, len(flags))
	for _, f := range flags {
		mapFlag[f] = true
	}

	return &Grep{pattern: pattern, flags: mapFlag, files: files}
}

func (g *Grep) Grep() []string {
	g.processFlags()

	matches := make([]match, 0)
	for _, fileName := range g.files {
		matches = append(matches, g.processFile(fileName)...)
	}

	return g.processOutput(matches)
}

func (g *Grep) processFile(fileName string) (matches []match) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	s := bufio.NewScanner(f)
	for i := 1; s.Scan(); i++ {

		isMatching := g.isLineMatching(s.Text())
		if g.isFlagActive(flagInvertedMatch) {
			isMatching = !isMatching // Invert the match
		}

		if isMatching {
			matches = append(matches, match{filename: fileName, line: i, text: s.Text()})

			if g.isFlagActive(flagPrintFilenameOnly) {
				// We only need to match one single line on a file, so if found, we can skip the
				// rest of the lines.
				return matches
			}
		}
	}

	return matches
}

func (g *Grep) isLineMatching(line string) bool {
	if g.isFlagActive(flagIgnoreCase) {
		line = strings.ToLower(line)
	}

	if g.isFlagActive(flagExactMatch) {
		return line == g.pattern

	}

	if strings.Index(line, g.pattern) != -1 {
		return true
	}
	return false
}

func (g *Grep) processFlags() {
	if g.isFlagActive(flagIgnoreCase) {
		g.pattern = strings.ToLower(g.pattern)
	}
}

func (g *Grep) processOutput(matches []match) (output []string) {
	for _, match := range matches {
		if g.isFlagActive(flagPrintFilenameOnly) {
			output = append(output, match.filename)
			continue
		}

		output = append(output, g.formatLine(match))
	}

	return output
}

func (g *Grep) formatLine(match match) string {
	filename := ""
	if len(g.files) > 1 {
		filename = fmt.Sprintf("%s:", match.filename)
	}

	if g.isFlagActive(flagPrintLineNumber) {
		return fmt.Sprintf("%s%d:%s", filename, match.line, match.text)
	}

	return fmt.Sprintf("%s%s", filename, match.text)
}

func (g *Grep) isFlagActive(flag string) bool {
	_, found := g.flags[flag]
	return found
}
