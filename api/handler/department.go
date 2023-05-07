package handler

import (
	"employee/api/presenter"
	"employee/pkg/department"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddDepartment(service department.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var reqBody department.CreateDeptParams
		err := c.BodyParser(&reqBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.DeptErrorResponse(err))
		}

		if reqBody.Name == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.DeptErrorResponse(errors.New("please specify department name")))
		}

		res, err := service.InsertDept(c.Context(), &reqBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.DeptErrorResponse(err))
		}

		return c.JSON(presenter.DeptSuccessResponse(&res))
	}
}

func GetDepartment(service department.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if id == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.DeptErrorResponse(errors.New("please specify department name")))
		}

		res, err := service.GetDeptById(c.Context(), id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.DeptErrorResponse(err))
		}

		return c.JSON(presenter.DeptSuccessResponse(&res))
	}
}
