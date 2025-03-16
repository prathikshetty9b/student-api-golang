package storage

import "github.com/prathikshetty9b/students-api/pkg/types"

type Storage interface {
	// Create a new student
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentByID(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudentById(id int64, name string, email string, age int) error
	DeleteStudentById(id int64) error
}
