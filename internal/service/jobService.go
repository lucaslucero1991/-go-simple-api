package service

import (
	"errors"
	"strings"
	"v0/internal/delivery/http/params"
	"v0/internal/delivery/http/request"
	"v0/internal/delivery/http/response"
	"v0/internal/entity"
	"v0/internal/repository"
)

type IJobService interface {
	CreateJob(jobRequest *request.JobRequest) (int, error)
	GetJob(jobParam *params.JobParam) ([]*response.JobResponse, error)
}

type JobService struct {
	jobRepository repository.IJobRepository
}

func NewJobService(jobRepository repository.IJobRepository) IJobService {
	return &JobService{jobRepository: jobRepository}
}

func (s *JobService) CreateJob(jobRequest *request.JobRequest) (int, error) {

	validator := NewJobValidatorService(jobRequest)
	err := validator.Validate()
	if err != nil {
		return 0, errors.New(err.Error())
	}

	entityJob := entity.NewJob(
		jobRequest.Name,
		jobRequest.Country,
		jobRequest.Salary,
		strings.Join(jobRequest.Skills, ","))

	jobID, err := s.jobRepository.CreateJob(entityJob)
	if err != nil {
		return 0, err
	}

	return jobID, nil
}

func (s *JobService) GetJob(jobParam *params.JobParam) ([]*response.JobResponse, error) {
	jobs, err := s.jobRepository.GetJobs(
		jobParam.Name,
		jobParam.Country,
		jobParam.SalaryMin,
		jobParam.SalaryMax)

	if err != nil {
		return nil, err
	}

	return response.ConvertToJobResponses(jobs), nil
}
