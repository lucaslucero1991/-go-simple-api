package params

import (
	"net/http"
	"strconv"
)

type JobParam struct {
	Name      string
	SalaryMin int
	SalaryMax int
	Country   string
}

func NewJobParam(r *http.Request) *JobParam {
	queryParams := r.URL.Query()
	jobParam := &JobParam{}

	if name := queryParams.Get("name"); name != "" {
		jobParam.Name = name
	}

	if salaryMinStr := queryParams.Get("salary_min"); salaryMinStr != "" {
		salaryMin, err := strconv.Atoi(salaryMinStr)
		if err == nil {
			jobParam.SalaryMin = salaryMin
		}
	}

	if salaryMaxStr := queryParams.Get("salary_max"); salaryMaxStr != "" {
		salaryMax, err := strconv.Atoi(salaryMaxStr)
		if err == nil {
			jobParam.SalaryMax = salaryMax
		}
	}

	if country := queryParams.Get("country"); country != "" {
		jobParam.Country = country
	}

	return jobParam
}
