package service_test

import (
	"errors"
	"testing"
	"v0/internal/delivery/http/params"
	"v0/internal/mocks"
	service2 "v0/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"v0/internal/delivery/http/request"
	"v0/internal/entity"
)

func TestCreateJob_WhenCreateJobWasSuccess_ThenReturnJobId(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockJobRepository)
	mockExternalAPIRepo := new(mocks.MockExternalAPIRepository)
	mockRepo.On("CreateJob", mock.Anything).Return(123, nil)
	mockJobRequest := &request.JobRequest{
		Name:    "Software Engineer",
		Salary:  5000,
		Country: "USA",
		Skills:  []string{"Go", "Java", "Python"},
	}

	service := service2.NewJobService(mockRepo, mockExternalAPIRepo)

	// Act
	jobID, err := service.CreateJob(mockJobRequest)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 123, jobID)
}

func TestCreateJob_WhenCreateJobFailed_ThenReturnErr(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockJobRepository)
	mockExternalAPIRepo := new(mocks.MockExternalAPIRepository)
	mockRepo.On("CreateJob", mock.Anything).Return(0,
		errors.New("dummy error"))
	mockJobRequest := &request.JobRequest{
		Name:    "Software Engineer",
		Salary:  5000,
		Country: "USA",
		Skills:  []string{"Go", "Java", "Python"},
	}

	service := service2.NewJobService(mockRepo, mockExternalAPIRepo)

	// Act
	jobID, err := service.CreateJob(mockJobRequest)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, 0, jobID)
}

func TestGetJobs_WhenGetJobsWasSuccess_ThenReturnJobs(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockJobRepository)
	mockExternalAPIRepo := new(mocks.MockExternalAPIRepository)
	expectedJobs := []*entity.Job{
		{ID: 1, Name: "Software Engineer", Salary: 5000, Country: "USA", Skills: "Go,Java,Python"},
		{ID: 2, Name: "Data Scientist", Salary: 6000, Country: "Canada", Skills: "Python,R"},
	}
	mockRepo.On("GetJobs", "", "", 0, 0).Return(expectedJobs, nil)
	mockExternalAPIRepo.On("GetJobs", "", "", 0, 0).Return(expectedJobs, nil)

	service := service2.NewJobService(mockRepo, mockExternalAPIRepo)
	jobParam := &params.JobParam{}

	// Act
	jobs, err := service.GetJob(jobParam)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, jobs)
	assert.Len(t, jobs, 4)
	assert.Equal(t, expectedJobs[0].Name, jobs[0].Name)
	assert.Equal(t, expectedJobs[0].Salary, jobs[0].Salary)
	assert.Equal(t, expectedJobs[0].Country, jobs[0].Country)
}

func TestGetJobs_WhenGetJobsFailed_ThenReturnErr(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockJobRepository)
	mockExternalAPIRepo := new(mocks.MockExternalAPIRepository)
	var nilEntity []*entity.Job
	mockRepo.On("GetJobs", "", "", 0, 0).
		Return(nilEntity, errors.New("external api error"))
	mockExternalAPIRepo.On("GetJobs", "", "", 0, 0).
		Return(nilEntity, errors.New("external api error"))

	service := service2.NewJobService(mockRepo, mockExternalAPIRepo)
	jobParam := &params.JobParam{}

	// Act
	jobs, err := service.GetJob(jobParam)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, jobs)
}
