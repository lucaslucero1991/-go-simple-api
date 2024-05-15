package useCase

import (
	"fmt"
	"v0/internal/domain"
)

const (
	ErrMissingNameField    = "missing name field in job request"
	ErrMissingCountryField = "missing country field in job request"
	ErrMissingSalaryField  = "missing salary field in job request"
	ErrMissingSkillsField  = "missing skills field in job request"
)

type IValidatorUseCase interface {
	Validate() error
}

type JobValidatorUseCase struct {
	jobDomain *domain.Job
}

func NewJobValidatorUseCase(jobDomain *domain.Job) IValidatorUseCase {
	return &JobValidatorUseCase{
		jobDomain: jobDomain,
	}
}

func (s *JobValidatorUseCase) Validate() error {
	if s.jobDomain.Name == "" {
		return fmt.Errorf(ErrMissingNameField)
	}
	if s.jobDomain.Country == "" {
		return fmt.Errorf(ErrMissingCountryField)
	}
	if s.jobDomain.Salary == 0 {

		return fmt.Errorf(ErrMissingSalaryField)
	}
	if len(s.jobDomain.Skills) == 0 {

		return fmt.Errorf(ErrMissingSkillsField)
	}
	return nil
}
