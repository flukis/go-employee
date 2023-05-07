package department

import (
	"context"
	"database/sql"
	"employee/helpers"
	"employee/pkg/entities"
	"errors"
	"regexp"
	"strconv"
)

type Service interface {
	InsertDept(context.Context, *CreateDeptParams) (entities.Department, error)
	GetDeptById(context.Context, string) (entities.Department, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Create Department
type CreateDeptParams struct {
	Name string `json:"name"`
}

func (s *service) InsertDept(ctx context.Context, arg *CreateDeptParams) (res entities.Department, err error) {
	_, err = s.repository.GetByName(ctx, arg.Name)
	if err != sql.ErrNoRows {
		err = errors.New("name already used, pick another name")
		return
	}

	lastDeptId, err := s.repository.GetLastDeptId(ctx)
	if err != nil {
		return
	}

	re := regexp.MustCompile("[0-9]+")
	idStr := re.FindAllString(lastDeptId, -1)
	id, err := strconv.Atoi(idStr[0])

	if err != nil {
		return
	}

	generatedId, err := helpers.GenerateId(id)

	if err != nil {
		return
	}

	return s.repository.Create(ctx, &entities.Department{
		DepartmentName: arg.Name,
		DepartmentNo:   generatedId,
	})
}

type GetDeptParams struct {
	ID string `json:"id"`
}

func (s *service) GetDeptById(ctx context.Context, id string) (res entities.Department, err error) {
	res, err = s.repository.GetById(ctx, id)
	return
}
