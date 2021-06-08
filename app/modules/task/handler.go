package task

import (
	"encoding/json"
	"net/http"
	"test/app/model"
	"test/app/repository"

	"github.com/gofiber/fiber/v2"
)

type handler struct{}

func Handler() handler {
	return handler{}
}

func (h handler) CreateTask(c *fiber.Ctx) error {
	service := NewService()
	task := model.Task{}
	c.BodyParser(&task)
	service.CreateTask(task.Sentence)
	message := []byte("ok")
	return c.Status(http.StatusOK).Send(message)
}

func (h handler) ListTask(c *fiber.Ctx) error {
	service := NewService()
	list, _ := json.Marshal(service.TaskList())
	message := []byte(string(list))
	return c.Status(http.StatusOK).Send(message)
}

var NewService = func() *Service {
	repo := repository.NewTaskDB()
	return &Service{Repo: repo}
}
