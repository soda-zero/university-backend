package models

import "time"

type CareerEnrollment struct {
	ID             string    `json:"id"`
	StudentID      string    `json:"student_id"`
	CareerID       string    `json:"career_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`
	CareerStatusID string    `json:"career_status_id"`
}
