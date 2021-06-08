package task_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"test/app/model"
	"test/app/modules/task"
	"test/server"
	"test/test/mocks"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("ENVIRONMENT", "testing")
	task.NewService = func() *task.Service {
		return &task.Service{mocks.NewMockTaskDB()}
	}
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCreateTask(t *testing.T) {
	app := server.RunServer()

	task := model.Task{Sentence: "La Prueba Para El Test"}
	newJson, _ := json.Marshal(task)
	req := httptest.NewRequest("POST", "/test/task", bytes.NewReader(newJson))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	if err != nil {
		t.Fail()
	}

	if res.StatusCode != 200 {
		t.Fail()
	}

	body, _ := ioutil.ReadAll(res.Body)
	if string(body) != "ok" {
		t.Fail()
	}
}

func TestListTask(t *testing.T) {
	app := server.RunServer()

	req := httptest.NewRequest("GET", "/test/task", nil)
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	if err != nil {
		t.Fail()
	}

	if res.StatusCode != 200 {
		t.Fail()
	}

	list, _ := json.Marshal(mocks.NewMockTaskDB().ListTask())
	body, _ := ioutil.ReadAll(res.Body)
	if string(body) != string(list) {
		t.Fail()
	}
}
