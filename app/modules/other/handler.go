package other

import (
	"encoding/json"
	"net/http"
	"test/app/model"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service Service
}

func Handler(service Service) handler {
	return handler{service}
}

func (h handler) CreateTask(c *fiber.Ctx) error {
	task := model.Task{}
	c.BodyParser(&task)
	h.service.CreateTask(task.Sentence)
	message := []byte("ok")
	return c.Status(http.StatusOK).Send(message)
}

func (h handler) ListTask(c *fiber.Ctx) error {
	list, _ := json.Marshal(h.service.TaskList())
	message := []byte(string(list))
	return c.Status(http.StatusOK).Send(message)
}
