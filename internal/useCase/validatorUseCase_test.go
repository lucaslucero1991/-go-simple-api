package useCase_test

import (
	"errors"
	"testing"
	"v0/internal/domain"
	"v0/internal/useCase"

	"github.com/stretchr/testify/assert"
	"v0/internal/delivery/http/request"
)

func TestJobValidatorUseCase_Validate(t *testing.T) {
	tests := []struct {
		name        string
		jobRequest  *request.JobRequest
		expectedErr error
	}{
		{
			name: "Valid job request",
			jobRequest: &request.JobRequest{
				Name:    "Software Engineer",
				Salary:  5000,
				Country: "USA",
				Skills:  []string{"Go", "Java", "Python"},
			},
			expectedErr: nil,
		},
		{
			name: "Missing name field",
			jobRequest: &request.JobRequest{
				Salary:  5000,
				Country: "USA",
				Skills:  []string{"Go", "Java", "Python"},
			},
			expectedErr: errors.New(useCase.ErrMissingNameField),
		},
		{
			name: "Missing country field",
			jobRequest: &request.JobRequest{
				Name:   "Software Engineer",
				Salary: 5000,
				Skills: []string{"Go", "Java", "Python"},
			},
			expectedErr: errors.New(useCase.ErrMissingCountryField),
		},
		{
			name: "Missing salary field",
			jobRequest: &request.JobRequest{
				Name:    "Software Engineer",
				Country: "USA",
				Skills:  []string{"Go", "Java", "Python"},
			},
			expectedErr: errors.New(useCase.ErrMissingSalaryField),
		},
		{
			name: "Missing skills field",
			jobRequest: &request.JobRequest{
				Name:    "Software Engineer",
				Salary:  5000,
				Country: "USA",
			},
			expectedErr: errors.New(useCase.ErrMissingSkillsField),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jobDomain := domain.NewDomainJob(test.jobRequest)
			validator := useCase.NewJobValidatorUseCase(jobDomain)
			err := validator.Validate()

			assert.Equal(t, test.expectedErr, err)
		})
	}
}
