package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	l := List{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, 3, l.Len())
	assert.Equal(t, "123", l.String())
	assert.Equal(t, 1, l.First())
	assert.Equal(t, 3, l.Last())
}

func TestPushFront(t *testing.T) {
	l := List{}
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	assert.Equal(t, 3, l.Len())
	assert.Equal(t, "321", l.String())
	assert.Equal(t, 3, l.First())
	assert.Equal(t, 1, l.Last())
}

func TestPushBackAndFront(t *testing.T) {
	l := List{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushFront(3)
	assert.Equal(t, 3, l.Len())
	assert.Equal(t, "312", l.String())
	assert.Equal(t, 3, l.First())
	assert.Equal(t, 2, l.Last())
}

func TestRemove(t *testing.T) {
	l := List{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.Remove(2)
	assert.Equal(t, 2, l.Len())
	assert.Equal(t, "13", l.String())
	assert.Equal(t, 1, l.First())
	assert.Equal(t, 3, l.Last())
}

func TestRemoveLastAndAdd(t *testing.T) {
	l := List{}

	l.PushBack(2)
	l.PushBack(3)
	l.PushFront(1)
	l.Remove(3)
	l.PushBack(4)
	assert.Equal(t, 3, l.Len())
	assert.Equal(t, "124", l.String())
	assert.Equal(t, 1, l.First())
	assert.Equal(t, 4, l.Last())
}

func TestRemoveFirstAndAdd(t *testing.T) {
	l := List{}

	l.PushBack(2)
	l.PushBack(3)
	l.PushFront(1)
	l.Remove(1)
	l.PushFront(0)
	assert.Equal(t, 3, l.Len())
	assert.Equal(t, "023", l.String())
	assert.Equal(t, 0, l.First())
	assert.Equal(t, 3, l.Last())
}

func TestNextPrev(t *testing.T) {
	l := List{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, 3, l.Len())
	assert.Equal(t, 3, l.Next(2))
	assert.Equal(t, 1, l.Prev(2))
}
func TestNextPrevString(t *testing.T) {
	l := List{}
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	assert.Equal(t, 3, l.Len())
	assert.Equal(t, "abc", l.String())
	assert.Equal(t, "a", l.First())
	assert.Equal(t, "c", l.Last())
}
