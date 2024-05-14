package service_test

import (
	"testing"
	service2 "v0/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"v0/internal/delivery/http/request"
	"v0/internal/entity"
)

type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) CreateJob(job *entity.Job) (int, error) {
	args := m.Called(job)
	return args.Int(0), args.Error(1)
}

func TestCreateJob(t *testing.T) {
	// Arrange
	mockRepo := new(MockJobRepository)
	mockRepo.On("CreateJob", mock.Anything).Return(123, nil)
	mockJobRequest := &request.JobRequest{
		Name:    "Software Engineer",
		Salary:  5000,
		Country: "USA",
		Skills:  []string{"Go", "Java", "Python"},
	}

	service := service2.NewJobService(mockRepo)

	// Act
	jobID, err := service.CreateJob(mockJobRequest)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 123, jobID)
}
