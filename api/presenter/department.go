package presenter

import (
	"employee/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type Department struct {
	ID   string `json:"dept_no"`
	Name string `json:"dept_name"`
}

func DeptSuccessResponse(data *entities.Department) *fiber.Map {
	d := Department{
		ID:   data.DepartmentNo,
		Name: data.DepartmentName,
	}
	return &fiber.Map{
		"status": true,
		"data":   d,
		"error":  nil,
	}
}

func DeptsSuccessResponse(data *[]Department) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func DeptErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  err.Error(),
	}
}
