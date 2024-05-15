package domain

import "v0/internal/delivery/http/request"

type Job struct {
	Name    string
	Country string
	Salary  int
	Skills  []string
}

func NewDomainJob(jobRequest *request.JobRequest) *Job {
	return &Job{
		Name:    jobRequest.Name,
		Country: jobRequest.Country,
		Salary:  jobRequest.Salary,
		Skills:  jobRequest.Skills,
	}
}
