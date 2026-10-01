package task

type Task struct {
	Title string
	Done  bool
}

func New(title string) Task {
	return Task{Title: title, Done: false}
}
