package types

import "fmt"

type Employee struct {
	Name         string `json:"name"`
	UserID       int    `json:"userId"`
	Age          int    `json:"age"`
	DepartmentID int    `json:"departmentId"`
}

func (e Employee) String() string {
	return fmt.Sprintf(
		"User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
		e.UserID,
		e.Age,
		e.Name,
		e.DepartmentID,
	)
}
