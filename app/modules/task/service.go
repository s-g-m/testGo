package task

import (
	"strings"
	"test/app/model"
	"test/app/tool"
)

type Service struct {
	Repo Repository
}

func (s *Service) CreateTask(sentence string) {
	list := strings.Split(sentence, " ")
	s.Repo.SaveTask(tool.CreateTaskA(list))
}

func (s *Service) TaskList() (tasks []model.Task) {
	return s.Repo.ListTask()
}
