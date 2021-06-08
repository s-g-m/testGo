package repository

import (
	"fmt"
	"sync"
	"test/app/model"
)

type TaskDB struct {
	count int
	tasks []model.Task
	mux   sync.Mutex
}

var taskDB *TaskDB
var once sync.Once

func NewTaskDB() *TaskDB {
	once.Do(func() {
		taskDB = &TaskDB{}
	})
	return taskDB
}

func (r *TaskDB) SaveTask(task model.Task) {
	r.mux.Lock()
	r.count++
	r.tasks = append(r.tasks, task)
	r.mux.Unlock()
}

func (r *TaskDB) ListTask() []model.Task {
	fmt.Println(r.count)
	return r.tasks
}
