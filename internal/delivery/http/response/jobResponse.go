package response

import (
	"strings"
	"v0/internal/entity"
)

type JobResponse struct {
	ID      int      `json:"Id"`
	Name    string   `json:"name"`
	Country string   `json:"country"`
	Salary  int      `json:"salary"`
	Skills  []string `json:"skills"`
}

func ConvertToJobResponses(jobs []*entity.Job) []*JobResponse {
	jobResponses := make([]*JobResponse, len(jobs))
	for i, job := range jobs {
		jobResponses[i] = &JobResponse{
			ID:      job.ID,
			Name:    job.Name,
			Country: job.Country,
			Salary:  job.Salary,
			Skills:  strings.Split(job.Skills, ","),
		}
	}
	return jobResponses
}
