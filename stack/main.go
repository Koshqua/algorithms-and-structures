//I decided to make a simple app with usage of stack.
//It will be a text block which can remember previous state and undo changes if needed
package main

import "fmt"

//TextBlock is out implementation of text block with cached changes
type TextBlock struct {
	currentText string
	cache       Stack
}

//NewTextBlock returns a new text block
func NewTextBlock(text string) *TextBlock {
	t := &TextBlock{currentText: text, cache: Stack{}}
	return t
}

//ChangeText changes text block text and adds previous to the cache
func (t *TextBlock) ChangeText(newText string) {
	t.cache.Push(t.currentText)
	t.currentText = newText
}

//Undo changes text of block to the previous one
func (t *TextBlock) Undo() {
	if t.cache.Len() > 0 {
		t.currentText = t.cache.Pop()
		return
	}
	fmt.Println("Cache is empty")
}

func main() {
	t := NewTextBlock("Initial text")
	t.ChangeText("This is first change")
	t.ChangeText("This is second change")
	t.ChangeText("This is third change")
	fmt.Println(t.currentText)
	fmt.Println("Undo once")
	t.Undo()
	fmt.Println(t.currentText)
	fmt.Println("Undo to the begining")
	t.Undo()
	t.Undo()
	fmt.Println(t.currentText)
	fmt.Println("Undo one more time - nothing would happend")
	t.Undo()
	fmt.Println(t.currentText)
}
