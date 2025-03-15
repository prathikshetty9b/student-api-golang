package sqllite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/prathikshetty9b/students-api/pkg/config"
	"github.com/prathikshetty9b/students-api/pkg/types"
)

type Sqllite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqllite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	name TEXT, 
	age INTEGER, 
	email TEXT)`)

	if err != nil {
		return nil, err
	}

	return &Sqllite{Db: db}, nil
}

func (s *Sqllite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Sqllite) GetStudentByID(id int64) (types.Student, error) {
	var student types.Student
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return student, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student not found with id %s: %w", fmt.Sprint(id), err)
		}
		return types.Student{}, fmt.Errorf("query failed: %w", err)
	}
	return student, nil
}
