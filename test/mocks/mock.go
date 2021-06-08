package mocks

import (
	"test/app/model"
)

type MockTaskDB struct {
	tasks []model.Task
}

func NewMockTaskDB() MockTaskDB {
	mapa := make(map[string]int)
	mapa["l"] = 2
	mapa["_"] = 4
	mapa["a"] = 4
	mapa["b"] = 1
	mapa["e"] = 3
	mapa["p"] = 2
	mapa["r"] = 2
	mapa["s"] = 1
	mapa["t"] = 2
	mapa["u"] = 1
	tasks := []model.Task{{Values: mapa, Sentence: "la_prueba_para_el_test"}, {Values: mapa, Sentence: "la_prueba_para_el_test"}}
	return MockTaskDB{tasks: tasks}
}

func (r MockTaskDB) SaveTask(task model.Task) {

}

func (r MockTaskDB) ListTask() []model.Task {
	return r.tasks
}
