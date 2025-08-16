package usecase

import "grpc_test/internal/domain"

type UseCase struct {
	repo domain.TodoList
}

func NewUseCase(repo domain.TodoList) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) Create(title, description string) domain.Task {
	return u.repo.CreateTask(title, description)
}

func (u *UseCase) GetList() domain.TasksList {
	return u.repo.ListTasks()
}

func (u *UseCase) Edit(taskRequest domain.Task) domain.Task {
	return u.repo.EditTask(taskRequest)
}

func (u *UseCase) Delete(id int) domain.Task {
	return u.repo.DeleteTask(id)
}
