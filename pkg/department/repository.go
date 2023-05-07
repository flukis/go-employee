package department

import (
	"context"
	"employee/pkg/entities"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Create(context.Context, *entities.Department) (entities.Department, error)
	GetLastDeptId(context.Context) (string, error)
	GetByName(context.Context, string) (string, error)
}

type postgreRepository struct {
	db *sqlx.DB
}

func NewPotgreRepository(db *sqlx.DB) Repository {
	return &postgreRepository{
		db: db,
	}
}

func (p *postgreRepository) Create(ctx context.Context, arg *entities.Department) (res entities.Department, err error) {
	queryStr := `
		INSERT INTO departments (
			dept_no,
			dept_name
		) VALUES (
			$1,
			$2
		) RETURNING *
	`
	res = entities.Department{}
	err = p.db.QueryRowContext(ctx, queryStr, arg.DepartmentNo, arg.DepartmentName).Scan(&res.DepartmentNo, &res.DepartmentName)
	return
}

func (p *postgreRepository) GetLastDeptId(ctx context.Context) (res string, err error) {
	dept := entities.Department{}
	queryStr := `
		SELECT dept_no, dept_name FROM departments
		ORDER BY dept_no DESC
		LIMIT 1
	`
	err = p.db.GetContext(ctx, &dept, queryStr)
	res = dept.DepartmentNo
	return
}

func (p *postgreRepository) GetByName(ctx context.Context, name string) (res string, err error) {
	dept := entities.Department{}
	queryStr := `
		SELECT dept_no, dept_name FROM departments
		WHERE dept_name = $1
	`
	err = p.db.GetContext(ctx, &dept, queryStr, name)
	res = dept.DepartmentNo
	return
}
