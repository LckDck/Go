package list

//Item is an item from the list
type Item struct {
	Value interface{}
	Prev  *Item
	Next  *Item
}

//List is double linked list
type List struct {
	len         int
	lastElement Item
}

//Len returns the length of the list
func (l List) Len() int {
	return l.len
}

//First returns the first element of the list
func (l List) First() interface{} {
	first, found := l.firstItem()
	if found {
		return first.Value
	}
	return 0
}

//Last returns the last element of the list
func (l List) Last() interface{} {
	if l.len == 0 {
		return 0
	}
	return l.lastElement.Value
}

//PushFront adds an element to the beginning
func (l List) PushFront(element interface{}) {
	first, found := l.firstItem()
	item := Item{Value: element, Prev: nil, Next: nil}
	if found {
		item.Next = &first
		first.Prev = &item
	} else {
		l.lastElement = item
	}
	l.len++
}

//PushBack adds an element to the end
func (l List) PushBack(element interface{}) {
	l.len++
	l.lastElement = Item{Value: element, Next: nil, Prev: l.lastElement.Prev}
}

//Remove removes an element from the list
func (l List) Remove(element interface{}) {
	item, found := l.find(element)
	if found {
		item.Prev.Next = item.Next
		item.Next.Prev = item.Prev
		l.len = l.len - 1
	}
}

//Prev returns the previous element from @element
func (l List) Prev(element *interface{}) interface{} {
	item, found := l.find(element)
	if found {
		return item.Prev.Value
	}
	return nil
}

//Next returns the last element of the list
func (l List) Next(element *interface{}) interface{} {
	item, found := l.find(element)
	if found {
		return item.Next.Value
	}
	return nil
}

func (l List) find(element interface{}) (Item, bool) {
	for curr := l.lastElement; curr.Prev != nil; curr = *curr.Prev {
		if element == curr.Value {
			return curr, false
		}
	}
	return Item{}, true
}

func (l List) firstItem() (Item, bool) {
	if l.len == 0 {
		return Item{}, false
	}
	for curr := l.lastElement; ; curr = *curr.Prev {
		if curr.Prev == nil {
			return curr, true
		}
	}
}

// элемент списка Value() interface{} ??
