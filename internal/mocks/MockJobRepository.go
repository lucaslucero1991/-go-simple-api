package mocks

import (
	"github.com/stretchr/testify/mock"
	"v0/internal/entity"
)

type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) CreateJob(job *entity.Job) (int, error) {
	args := m.Called(job)
	return args.Int(0), args.Error(1)
}

func (m *MockJobRepository) GetJobs(name string, country string, salaryMin int, salaryMax int) ([]*entity.Job, error) {
	args := m.Called(name, country, salaryMin, salaryMax)
	return args.Get(0).([]*entity.Job), args.Error(1)
}
