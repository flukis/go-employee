package entities

// type Employee struct {
// 	EmployeeNo int       `json:"emp_no"`
// 	BirthDate  time.Time `json:"birth_date"`
// 	HireDate   time.Time `json:"hire_date"`
// 	Gender     string    `json:"gender"`
// 	FirstName  string    `json:"first_name"`
// 	LastName   string    `json:"last_name"`
// }

// type DepartmentEmployee struct {
// 	EmployeeNo   string    `json:"emp_no"`
// 	DepartmentNo int       `json:"dept_no"`
// 	FromDate     time.Time `json:"from_date"`
// 	ToDate       time.Time `json:"to_date"`
// }

// type Salary struct {
// 	EmployeeNo int       `json:"emp_no"`
// 	Salary     int       `json:"salary"`
// 	FromDate   time.Time `json:"from_date"`
// 	ToDate     time.Time `json:"to_date"`
// }

// type Title struct {
// 	EmployeeNo int       `json:"emp_no"`
// 	Title      string    `json:"title"`
// 	FromDate   time.Time `json:"from_date"`
// 	ToDate     time.Time `json:"to_date"`
// }

// type DepartmentManager struct {
// 	EmployeeNo   string    `json:"emp_no"`
// 	DepartmentNo int       `json:"dept_no"`
// 	FromDate     time.Time `json:"from_date"`
// 	ToDate       time.Time `json:"to_date"`
// }

type Department struct {
	DepartmentName string `json:"dept_name" db:"dept_name"`
	DepartmentNo   string `json:"dept_no" db:"dept_no"`
}
