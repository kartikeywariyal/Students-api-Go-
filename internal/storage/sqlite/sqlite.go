package sqlite

import (
	"database/sql"
	"math/rand"

	"github.com/kartikeywariyal/students-api-Go-/internal/config"
	"github.com/kartikeywariyal/students-api-Go-/internal/types"
)

type SqliteStorage struct {
	db *sql.DB
}

func (s *SqliteStorage) CreateStudent(name string, age string, email string) (int64, error) {
	_, err := s.db.Exec(`Insert into students (name,age,email) values (?,?,?)`, name, age, email)
	if err != nil {
		return 0, err
	}

	return rand.Int63(), err
}

func (s *SqliteStorage) GetStudent(id int64) (types.Student, error) {
	row := s.db.QueryRow(`Select id,name,age,email from students where id=?`, id)
	var student types.Student
	err := row.Scan(&student.ID, &student.Name, &student.Age, &student.Email)
	if err != nil {
		return types.Student{}, err
	}
	return student, nil
}

func NewSqliteStorage(cfg *config.Config) (sq *SqliteStorage, err error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (id INTEGER PRIMARY KEY AUTOINCREMENT, name varchar(30), age INTEGER, email TEXT)`)
	if err != nil {
		return nil, err
	}
	return &SqliteStorage{db: db}, nil
}
