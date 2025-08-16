package domain

type Task struct {
	ID          int64
	Title       string
	Description string
	Completed   bool
}

type TasksList struct {
	Tasks []Task
}

type TodoList interface {
	CreateTask(title, description string) Task
	ListTasks() TasksList
	EditTask(taskRequest Task) Task
	DeleteTask(id int) Task
}
