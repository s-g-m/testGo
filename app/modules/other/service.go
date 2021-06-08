package other

import (
	"strings"
	"test/app/model"
	"test/app/tool"
)

type Service interface {
	CreateTask(sentence string)
	TaskList() (tasks []model.Task)
}

type service struct {
	Repo Repository
}

func NewService(repo Repository) service {
	return service{Repo: repo}
}

func (s service) CreateTask(sentence string) {
	list := strings.Split(sentence, " ")
	s.Repo.SaveTask(tool.CreateTaskA(list))
}

func (s service) TaskList() (tasks []model.Task) {
	return s.Repo.ListTask()
}
