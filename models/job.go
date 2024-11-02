// models/job.go
package models

import (
	"syndio-backend-app/database"
)

type Job struct {
	EmployeeID int    `json:"employee_id"`
	Department string `json:"department"`
	JobTitle   string `json:"job_title"`
}

func (job *Job) Save() error {
	_, err := database.DB.Exec("INSERT INTO jobs (employee_id, department, job_title) VALUES (?, ?, ?)", job.EmployeeID, job.Department, job.JobTitle)
	return err
}
