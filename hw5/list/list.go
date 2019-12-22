package list

import "fmt"

//Item is an item from the list
type Item struct {
	Value interface{}
	Prev  *Item
	Next  *Item
}

//List is double linked list
type List struct {
	len          int
	lastElement  *Item
	firstElement *Item
}

//Len returns the length of the list
func (l List) Len() int {
	return l.len
}

//First returns the first element of the list
func (l *List) First() interface{} {
	if l.len == 0 {
		return 0
	}
	return l.firstElement.Value
}

//Last returns the last element of the list
func (l *List) Last() interface{} {
	if l.len == 0 {
		return 0
	}
	return l.lastElement.Value
}

//PushFront adds an element to the beginning
func (l *List) PushFront(element interface{}) {
	item := Item{Value: element, Prev: nil, Next: nil}
	if l.len > 0 {
		item.Next = l.firstElement
		l.firstElement.Prev = &item
		l.firstElement = &item
	} else {
		l.lastElement = &item
		l.firstElement = &item
	}
	l.len++
}

//PushBack adds an element to the end
func (l *List) PushBack(element interface{}) {
	item := Item{Value: element, Prev: nil, Next: nil}
	if l.len > 0 {
		item.Prev = l.lastElement
		l.lastElement.Next = &item
		l.lastElement = &item
	} else {
		l.lastElement = &item
		l.firstElement = &item
	}
	l.len++
}

//Remove removes an element from the list
func (l *List) Remove(element interface{}) {
	item, found := l.find(element)
	if !found {
		return
	}
	if l.Len() == 1 {
		l.firstElement = nil
		l.lastElement = nil
	}
	if item.Prev != nil {
		item.Prev.Next = item.Next
	} else {
		l.firstElement = item.Next
		l.firstElement.Prev = nil
	}
	if item.Next != nil {
		item.Next.Prev = item.Prev
	} else {
		l.lastElement = item.Prev
		l.lastElement.Next = nil
	}
	l.len = l.len - 1
}

//Prev returns the previous element from @element
func (l List) Prev(element interface{}) interface{} {
	item, found := l.find(element)
	if found {
		return item.Prev.Value
	}
	return nil
}

//Next returns the last element of the list
func (l List) Next(element interface{}) interface{} {
	item, found := l.find(element)
	if found {
		return item.Next.Value
	}
	return nil
}

//String returns string representation of the list
func (l *List) String() string {
	var s = ""
	for curr := l.lastElement; ; curr = curr.Prev {
		s = fmt.Sprintf("%v", curr.Value) + s
		if curr.Prev == nil {
			return s
		}
	}
}

func (l List) find(element interface{}) (*Item, bool) {
	for curr := l.lastElement; curr != nil; curr = curr.Prev {
		if element == curr.Value {
			return curr, true
		}
	}
	return nil, false
}

// элемент списка Value() interface{} ??
