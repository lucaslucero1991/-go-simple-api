package service

import (
	"strings"
	"v0/internal/delivery/http/request"
	"v0/internal/entity"
	"v0/internal/repository"
)

type IJobService interface {
	CreateJob(jobRequest *request.JobRequest) (int, error)
}

type JobService struct {
	jobRepository repository.IJobRepository
}

func NewJobService(jobRepository repository.IJobRepository) IJobService {
	return &JobService{jobRepository: jobRepository}
}

func (s *JobService) CreateJob(jobRequest *request.JobRequest) (int, error) {
	entityJob := &entity.Job{
		Name:    jobRequest.Name,
		Salary:  jobRequest.Salary,
		Country: jobRequest.Country,
		Skills:  strings.Join(jobRequest.Skills, ","),
	}

	jobID, err := s.jobRepository.CreateJob(entityJob)
	if err != nil {
		return 0, err
	}

	return jobID, nil
}
