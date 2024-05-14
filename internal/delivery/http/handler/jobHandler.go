package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"v0/internal/delivery/http/request"
	"v0/internal/service"
)

type JobHandler struct {
	jobService service.IJobService
	validator  service.IValidatorService
}

func NewJobHandler(jobService service.IJobService,
	validator service.IValidatorService) *JobHandler {
	return &JobHandler{
		jobService: jobService,
		validator:  validator}
}

func (j *JobHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var jobRequest request.JobRequest
	err := json.NewDecoder(r.Body).Decode(&jobRequest)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	err = j.validator.Validate(&jobRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jobId, err := j.jobService.CreateJob(&jobRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/jobs/%d", jobId))
	w.WriteHeader(http.StatusCreated)
}
