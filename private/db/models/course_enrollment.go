package models

import "database/sql"

type CourseEnrollment struct {
	ID                 string          `json:"id"`
	StudentID          string          `json:"student_id"`
	CourseOccurrenceID string          `json:"course_occurrence_id"`
	FinalScore         sql.NullFloat64 `json:"final_score"`
}
