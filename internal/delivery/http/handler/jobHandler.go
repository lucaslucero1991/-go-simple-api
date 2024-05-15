package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"v0/internal/delivery/http/params"
	"v0/internal/delivery/http/request"
	"v0/internal/useCase"
)

type JobHandler struct {
	jobUseCase useCase.IJobUseCase
}

func NewJobHandler(jobUseCase useCase.IJobUseCase) *JobHandler {
	return &JobHandler{
		jobUseCase: jobUseCase}
}

func (j *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var jobRequest request.JobRequest
	err := json.NewDecoder(r.Body).Decode(&jobRequest)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	jobId, err := j.jobUseCase.CreateJob(&jobRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/jobs/%d", jobId))
	w.WriteHeader(http.StatusCreated)
}

func (j *JobHandler) GetJob(w http.ResponseWriter, r *http.Request) {
	var jobRequest request.JobRequest
	err := jobRequest.ValidateQueryParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jobs, err := j.jobUseCase.GetJob(params.NewJobParam(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(jobs)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(responseBody)
	w.WriteHeader(http.StatusOK)
}
