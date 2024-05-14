package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"v0/internal/entity"
)

type IJobRepository interface {
	CreateJob(job *entity.Job) (int, error)
	GetJobs(name string, country string, salaryMin int, salaryMax int) ([]*entity.Job, error)
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

func (r *SQLiteJobRepository) GetJobs(name string,
	country string,
	salaryMin int,
	salaryMax int) ([]*entity.Job, error) {
	var jobs []*entity.Job

	query := "SELECT * FROM jobs WHERE 1=1"
	args := []interface{}{}

	if name != "" {
		query += " AND name = ?"
		args = append(args, name)
	}
	if country != "" {
		query += " AND country = ?"
		args = append(args, country)
	}
	if salaryMin > 0 {
		query += " AND salary >= ?"
		args = append(args, salaryMin)
	}
	if salaryMax > 0 {
		query += " AND salary <= ?"
		args = append(args, salaryMax)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var job entity.Job
		err := rows.Scan(&job.ID, &job.Name, &job.Salary, &job.Country, &job.Skills)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		jobs = append(jobs, &job)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %v", err)
	}

	return jobs, nil
}
