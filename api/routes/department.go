package routes

import (
	"employee/api/handler"
	"employee/pkg/department"

	"github.com/gofiber/fiber/v2"
)

func DeptRouter(app fiber.Router, service department.Service) {
	app.Post("/department", handler.AddDepartment(service))
}
