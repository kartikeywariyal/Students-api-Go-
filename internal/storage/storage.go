package storage

import "github.com/kartikeywariyal/students-api-Go-/internal/types"

type Storage interface {
	CreateStudent(name string, age string, email string) (int64, error)
	GetStudent(id int64) (types.Student, error)
}
