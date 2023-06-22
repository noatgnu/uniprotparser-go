package parser

import (
	"net/http"
)

type Job struct {
	JobURL       string
	session      *http.Client
	pollInterval int
}

type JobIDResponse struct {
	JobID string `json:"jobId"`
}

func NewJob(jobURL string, session *http.Client, pollInterval int) *Job {
	return &Job{jobURL, session, pollInterval}
}

func (j *Job) CheckJob() string {
	jobStats, err := j.session.Get(j.JobURL)
	if err != nil {
		panic(err)
	}
	defer jobStats.Body.Close()
	if jobStats.StatusCode == 303 {
		return jobStats.Header.Get("Location")
	}
	return ""
}
