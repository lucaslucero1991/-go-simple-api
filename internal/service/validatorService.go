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
	Validate() error
}

type JobValidatorService struct {
	jobRequest *request.JobRequest
}

func NewJobValidatorService(jobRequest *request.JobRequest) IValidatorService {
	return &JobValidatorService{jobRequest: jobRequest}
}

func (s *JobValidatorService) Validate() error {
	if s.jobRequest.Name == "" {
		return fmt.Errorf(ErrMissingNameField)
	}
	if s.jobRequest.Country == "" {
		return fmt.Errorf(ErrMissingCountryField)
	}
	if s.jobRequest.Salary == 0 {

		return fmt.Errorf(ErrMissingSalaryField)
	}
	if len(s.jobRequest.Skills) == 0 {

		return fmt.Errorf(ErrMissingSkillsField)
	}
	return nil
}
