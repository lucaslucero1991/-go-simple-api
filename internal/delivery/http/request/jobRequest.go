package request

import (
	"errors"
	"net/http"
)

type JobRequest struct {
	Name    string   `json:"name"`
	Country string   `json:"country"`
	Salary  int      `json:"salary"`
	Skills  []string `json:"skills"`
}

func (req *JobRequest) ValidateQueryParams(r *http.Request) error {
	queryParams := r.URL.Query()

	validParams := map[string]bool{
		"name":       true,
		"salary_min": true,
		"salary_max": true,
		"country":    true,
	}

	for param := range queryParams {
		if !validParams[param] {
			return errors.New("invalid query param")
		}
	}
	return nil
}
