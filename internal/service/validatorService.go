package service

import (
	"fmt"
	"v0/internal/delivery/http/request"
)

type IValidatorService interface {
	Validate(jobRequest *request.JobRequest) error
}

type JobValidatorService struct {
}

func NewJobValidatorService() IValidatorService {
	return &JobValidatorService{}
}

func (s *JobValidatorService) Validate(jobRequest *request.JobRequest) error {
	if jobRequest.Name == "" {
		return fmt.Errorf("missing name field in job request")
	}
	if jobRequest.Country == "" {
		return fmt.Errorf("missing country field in job request")
	}
	if jobRequest.Salary == 0 {
		return fmt.Errorf("missing salary field in job request")
	}
	if len(jobRequest.Skills) == 0 {
		return fmt.Errorf("missing skills field in job request")
	}
	return nil
}
