package other_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"test/app/model"
	"test/app/modules/other"
	"test/test/mocks"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var app *fiber.App

func initServer() {
	app = fiber.New()
	app.Use(logger.New())
	service := other.NewService(mocks.NewMockTaskDB())
	other := other.Handler(service)
	app.Post("test/other", other.CreateTask)
	app.Get("test/other", other.ListTask)
}

func TestMain(m *testing.M) {
	os.Setenv("ENVIRONMENT", "testing")
	initServer()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCreateTask(t *testing.T) {
	task := model.Task{Sentence: "La Prueba Para El Test"}
	newJson, _ := json.Marshal(task)
	req := httptest.NewRequest("POST", "/test/other", bytes.NewReader(newJson))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)
	if err != nil {
		t.Fail()
	}

	if res.StatusCode != 200 {
		t.Fail()
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if string(body) != "ok" {
		t.Fail()
	}
}

func TestListTask(t *testing.T) {
	req := httptest.NewRequest("GET", "/test/other", nil)
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
