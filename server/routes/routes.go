package routes

import (
	"test/app/modules/other"
	"test/app/modules/task"
	"test/app/repository"

	"github.com/gofiber/fiber/v2"
)

func Load(app *fiber.App) {
	api := app.Group("/test")
	taskHandler(api)
	otherHandler(api)
}

func taskHandler(app fiber.Router) {
	task := task.Handler()
	app.Post("task", task.CreateTask)
	app.Get("task", task.ListTask)
}

func otherHandler(app fiber.Router) {
	service := other.NewService(repository.NewTaskDB())
	other := other.Handler(service)
	app.Post("other", other.CreateTask)
	app.Get("other", other.ListTask)
}
