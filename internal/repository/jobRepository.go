package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"v0/internal/entity"
)

type IJobRepository interface {
	CreateJob(job *entity.Job) (int, error)
}

type SQLiteJobRepository struct {
	db *sql.DB
}

func NewSQLiteJobRepository(db *sql.DB) *SQLiteJobRepository {
	return &SQLiteJobRepository{db: db}
}

func (r *SQLiteJobRepository) CreateJob(job *entity.Job) (int, error) {
	query := "INSERT INTO jobs (name, salary, country, skills) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, job.Name, job.Salary, job.Country, job.Skills)
	if err != nil {
		return 0, fmt.Errorf("failed to create job: %v", err)
	}

	jobID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get job ID: %v", err)
	}
	return int(jobID), nil
}
