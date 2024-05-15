package useCase

import (
	"errors"
	"fmt"
	"sync"
	"v0/internal/delivery/http/params"
	"v0/internal/delivery/http/request"
	"v0/internal/delivery/http/response"
	"v0/internal/domain"
	"v0/internal/entity"
	"v0/internal/repository"
)

type IJobUseCase interface {
	CreateJob(jobRequest *request.JobRequest) (int, error)
	GetJob(jobParam *params.JobParam) ([]*response.JobResponse, error)
}

type JobUseCase struct {
	jobRepository         repository.IJobRepository
	externalAPIRepository repository.IExternalAPIRepository
}

func NewJobUseCase(
	jobRepository repository.IJobRepository,
	externalAPIRepository repository.IExternalAPIRepository) IJobUseCase {
	return &JobUseCase{
		jobRepository:         jobRepository,
		externalAPIRepository: externalAPIRepository}
}

func (s *JobUseCase) CreateJob(jobRequest *request.JobRequest) (int, error) {

	jobDomain := domain.NewDomainJob(jobRequest)

	validator := NewJobValidatorUseCase(jobDomain)
	err := validator.Validate()
	if err != nil {
		return 0, errors.New(err.Error())
	}

	jobEntity := entity.NewJob(jobDomain)
	jobID, err := s.jobRepository.CreateJob(jobEntity)
	if err != nil {
		return 0, err
	}

	return jobID, nil
}

func (s *JobUseCase) GetJob(jobParam *params.JobParam) ([]*response.JobResponse, error) {

	var wg sync.WaitGroup
	wg.Add(2)

	jobChan := make(chan []*entity.Job, 2)
	errChan := make(chan error, 2)

	go func() {
		defer wg.Done()
		jobs, err := s.jobRepository.GetJobs(jobParam.Name, jobParam.Country, jobParam.SalaryMin, jobParam.SalaryMax)
		if err != nil {
			errChan <- err
			return
		}
		jobChan <- jobs
	}()

	go func() {
		defer wg.Done()
		jobs, err := s.externalAPIRepository.GetJobs(jobParam.Name, jobParam.Country, jobParam.SalaryMin, jobParam.SalaryMax)
		if err != nil {
			errChan <- err
			return
		}
		jobChan <- jobs
	}()

	go func() {
		wg.Wait()
		close(jobChan)
		close(errChan)
	}()

	var allJobs []*entity.Job
	var allErrors []error
	for j := range jobChan {
		allJobs = append(allJobs, j...)
	}
	for err := range errChan {
		allErrors = append(allErrors, err)
	}

	if len(allErrors) > 0 {
		return nil, fmt.Errorf("errors occurred: %v", allErrors)
	}

	return response.ConvertToJobResponses(allJobs), nil
}
