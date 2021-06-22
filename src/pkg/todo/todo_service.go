package todo

type ToDo interface {
	Create(text string, isDone bool)
}
