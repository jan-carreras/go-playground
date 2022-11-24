package editor

import (
	"container/list"
	"errors"
	"strings"
)

// 1.3.44 Text editor buffer. Develop a data type for a buffer in a text editor
// that implements the following API:
/**
insert(c rune) -> adds C in cursor position
delete() -> deleted and return the character at the cursor
left(int k) -> move cursor K positions to the left
right(int k) -> move the cursor  position to the right
size() -> chars in the buffer
*/

var ErrNothingToDelete = errors.New("nothing to delete")

type TextEditorBuffer interface {
	Insert(c rune)
	Delete() (rune, error)
	Left(k int)
	Right(k int)
	Size() int
}

type Buffer struct {
	list    list.List
	current *list.Element
}

func (b *Buffer) Insert(c rune) {
	if b.current == nil {
		b.current = b.list.PushFront(c)
		return
	}

	b.current = b.list.InsertAfter(c, b.current)
}

func (b *Buffer) Delete() (rune, error) {
	if b.current == nil {
		return 0, ErrNothingToDelete
	}

	toDelete := b.current
	b.current = toDelete.Prev()
	b.list.Remove(toDelete)

	return toDelete.Value.(rune), nil
}

func (b *Buffer) Left(k int) {
	for i := 0; i < k && b.current != nil; i++ {
		b.current = b.current.Prev()
	}
}

func (b *Buffer) Right(k int) {
	if b.current == nil {
		b.current = b.list.Front()
		k--
	}

	for i := 0; i < k && b.current.Next() != nil; i++ {
		b.current = b.current.Next()
	}
}

func (b *Buffer) Size() int {
	return b.list.Len()
}

func (b *Buffer) String() string {
	if b.list.Len() == 0 {
		return ""
	}

	buf := strings.Builder{}
	if b.current == nil {
		buf.WriteRune('|')
	}

	n := b.list.Front()
	for n != nil {
		buf.WriteRune(n.Value.(rune))
		if n == b.current {
			buf.WriteRune('|')
		}
		n = n.Next()
	}

	return buf.String()
}
