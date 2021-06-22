package postgres

import "fmt"

type ToDoImpl struct {
	Message string
}

func (t *ToDoImpl) Create(text string, isDone bool) {
	fmt.Printf("I have a message: %s\n", t.Message)
	panic("oops!")
}
