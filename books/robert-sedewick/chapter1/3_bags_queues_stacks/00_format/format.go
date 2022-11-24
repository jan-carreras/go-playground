package format

import (
	"container/list"
	"fmt"
	"strings"
)

func List(l list.List) string {
	b := strings.Builder{}

	n := l.Front()

	for n != nil {
		b.WriteString(fmt.Sprintf("%v", n.Value))
		if n != l.Back() {
			b.WriteString(" -> ")
		}

		n = n.Next()
	}

	return b.String()
}
