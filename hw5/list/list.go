package list

//List is double linked list
type List struct{}

//Len returns the length of the list
func (l List) Len() int {
	return 1
}

//First returns the first element of the list
func (l List) First() string {
	return ""
}

//Last returns the last element of the list
func (l List) Last() string {
	return ""
}

//PushFront adds an element to the beginning
func (l List) PushFront(element string) {

}

//PushBack adds an element to the end
func (l List) PushBack(element string) {

}

//Remove removes an element from the list
func (l List) Remove(element string) {

}

//Prev returns the previous element from @element
func (l List) Prev(element *string) string {
	return ""
}

//Next returns the last element of the list
func (l List) Next(element *string) string {
	return ""
}

// элемент списка Value() interface{} ??
