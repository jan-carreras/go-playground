package kindergarten

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var ErrUnknownPlant = errors.New("unknown plant")

type plant string

func newPlant(p string) (plant, error) {
	switch p {
	case "grass", "clover", "radishes", "violets":
		return plant(p), nil
	}

	return "", fmt.Errorf("%w: %q", ErrUnknownPlant, p)
}

func newPlantFromInitial(plantInitial uint8) (plant, error) {
	plants := map[uint8]string{'R': "radishes", 'C': "clover", 'G': "grass", 'V': "violets"}

	p, ok := plants[plantInitial]
	if !ok {
		return "", fmt.Errorf("unknown initial: %q", plantInitial)
	}

	return newPlant(p)
}

type space struct {
	plants []plant
}

type childrenToSpace map[string]int

type Garden struct {
	spaces   []space
	children childrenToSpace
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	lines := strings.Split(diagram, "\n")
	if len(lines) != 3 {
		return nil, errors.New("expected 3 lines as input")
	}

	lines = lines[1:] // The first line is always empty
	if len(lines[0])+len(lines[1]) != len(children)*4 {
		return nil, errors.New("not every students can have 4 plants")
	}

	g := &Garden{}
	for i := 0; i < len(lines[0]); i += 2 {
		read := []byte{
			lines[0][i], lines[0][i+1],
			lines[1][i], lines[1][i+1],
		}

		s := space{}
		for _, r := range read {
			if p, err := newPlantFromInitial(r); err != nil {
				return nil, err
			} else {
				s.plants = append(s.plants, p)
			}
		}

		g.spaces = append(g.spaces, s)
	}

	childrenCopy := make([]string, len(children))
	copy(childrenCopy, children)
	sort.Strings(childrenCopy)

	g.children = make(childrenToSpace, len(childrenCopy))
	for i, c := range childrenCopy {
		if _, ok := g.children[c]; ok {
			return nil, fmt.Errorf("duplicate name: %s", c)
		}
		g.children[c] = i
	}

	return g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	idx, ok := g.children[child]
	if !ok {
		return nil, false
	}

	rsp := make([]string, 0, 4)
	for _, p := range g.spaces[idx].plants {
		rsp = append(rsp, string(p))
	}

	return rsp, true
}
