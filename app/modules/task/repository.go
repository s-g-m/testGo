package task

import "test/app/model"

type Repository interface {
	SaveTask(task model.Task)
	ListTask() (list []model.Task)
}
