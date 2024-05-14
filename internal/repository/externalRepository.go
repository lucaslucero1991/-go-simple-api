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

	var jobs [][]interface{}
	err = json.NewDecoder(resp.Body).Decode(&jobs)
	if err != nil {
		return nil, err
	}

	var jobList []*entity.Job
	for _, jobData := range jobs {
		job := entity.NewJob(jobData[0].(string),
			jobData[2].(string),
			int(jobData[1].(float64)),
			joinStringSlice(jobData[3].([]interface{})))
		jobList = append(jobList, job)
	}

	return jobList, nil
}

func joinStringSlice(slice []interface{}) string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = v.(string)
	}
	return strings.Join(strSlice, ",")
}
