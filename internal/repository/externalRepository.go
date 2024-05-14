package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"v0/internal/entity"
)

type ExternalJob struct {
	Name    string   `json:"name"`
	Salary  int      `json:"salary"`
	Country string   `json:"country"`
	Skills  []string `json:"skills"`
}

type IExternalAPIRepository interface {
	GetJobs(name string, country string, salaryMin int, salaryMax int) ([]*entity.Job, error)
}

type ExternalAPIRepository struct {
}

func NewExternalAPIRepository() *ExternalAPIRepository {
	return &ExternalAPIRepository{}
}

func (r *ExternalAPIRepository) GetJobs(name string,
	country string,
	salaryMin int,
	salaryMax int) ([]*entity.Job, error) {

	baseURL := "http://localhost:8081/jobs?"

	if name != "" {
		baseURL += "&name=" + url.QueryEscape(name)
	}
	if country != "" {
		baseURL += "&country=" + url.QueryEscape(country)
	}
	if salaryMin > 0 {
		baseURL += "&salary_min=" + strconv.Itoa(salaryMin)
	}
	if salaryMax > 0 {
		baseURL += "&salary_max=" + strconv.Itoa(salaryMax)
	}

	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("Error calling HTTP:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var externalJobs [][]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&externalJobs); err != nil {
		return nil, err
	}

	var jobs []*entity.Job
	for _, jobData := range externalJobs {
		job := &entity.Job{
			Name:    jobData[0].(string),
			Salary:  int(jobData[1].(float64)),
			Country: jobData[2].(string),
			Skills:  strings.Join(jobData[3].([]string), ", "),
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func joinStringSlice(slice []interface{}) string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = v.(string)
	}
	return strings.Join(strSlice, ",")
}
