package robotname

import (
	"errors"
	"fmt"
)

/**
Approach 1:

- Have a hashmap in global state (!!!) that stores the names being "served"
- Memory space: O(n)
- When we want to issue a new name we generate a random one and check if it exists on the map
	- The problem is that the more names we give, the more difficult is not to "randomly generate one already assigned name"

Approach 2:
- Brute force
- Generate all possible names and put them on a list of available names
- Every time we assign a name we remove it from the list
- We pick random items form the list until the list is empty
- Memory space O(n)

Have a map (like a set) with all the elements, in memory. Remove from map every time we assign a new name.

Approach 3:
- Create a hash function that maps an ever increasing number (int) to a deterministic robot name.
	- Each number will have only one representation as robot name
	- Having this hash function makes the problem trivial
*/

type AllNames struct {
	names map[string]interface{}
}

func NewAllNames() *AllNames {
	names := make(map[string]interface{})

	for x := 'A'; x <= 'Z'; x++ {
		for y := 'A'; y <= 'Z'; y++ {
			for z := 0; z < 1000; z++ {
				names[fmt.Sprintf("%c%c%03d", x, y, z)] = struct{}{}
			}
		}
	}

	return &AllNames{names: names}
}

func (n *AllNames) ReserveName() (string, error) {
	for name := range n.names {
		delete(n.names, name)
		return name, nil
	}

	return "", errors.New("no names available")
}

var names = NewAllNames()

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if newName, err := names.ReserveName(); err != nil {
		return "", err
	} else {
		r.name = newName
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
