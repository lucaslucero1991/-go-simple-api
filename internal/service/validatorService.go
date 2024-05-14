package service

import (
	"fmt"
	"v0/internal/delivery/http/request"
)

const (
	ErrMissingNameField    = "missing name field in job request"
	ErrMissingCountryField = "missing country field in job request"
	ErrMissingSalaryField  = "missing salary field in job request"
	ErrMissingSkillsField  = "missing skills field in job request"
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
		return fmt.Errorf(ErrMissingNameField)
	}
	if jobRequest.Country == "" {
		return fmt.Errorf(ErrMissingCountryField)
	}
	if jobRequest.Salary == 0 {

		return fmt.Errorf(ErrMissingSalaryField)
	}
	if len(jobRequest.Skills) == 0 {

		return fmt.Errorf(ErrMissingSkillsField)
	}
	return nil
}
