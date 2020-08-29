package main

//Stack represents our stack implementation
type Stack struct {
	top  *Element
	size int
}

//Element represents element of a stack
type Element struct {
	value string
	next  *Element
}

//Len returns stack current length
func (s *Stack) Len() int {
	return s.size
}

//Push adds element to the stack and increases it's size
func (s *Stack) Push(element string) {
	s.top = &Element{value: element, next: s.top}
	s.size++
}

//Pop removes top element from the stack and returns it's value
func (s *Stack) Pop() (value string) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return value
	}
	return ""
}

//Peek returns the value from the top of stack
func (s *Stack) Peek() string {
	if s.size > 0 {
		return s.top.value
	}
	return ""
}
