package entity

import (
	"strings"
	"v0/internal/domain"
)

type Job struct {
	ID      int
	Name    string
	Country string
	Salary  int
	Skills  string
}

func NewJob(jobRequest *domain.Job) *Job {
	return &Job{
		Name:    jobRequest.Name,
		Country: jobRequest.Country,
		Salary:  jobRequest.Salary,
		Skills:  strings.Join(jobRequest.Skills, ","),
	}
}

func NewJobEntity(name string, country string, salary int, skills string) *Job {
	return &Job{
		Name:    name,
		Country: country,
		Salary:  salary,
		Skills:  skills,
	}
}
