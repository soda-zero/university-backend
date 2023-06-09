package models

import (
	"database/sql"
)

type CareerEnrollmentEvaluation struct {
	ID                 string          `json:"id"`
	CourseEnrollmentID string          `json:"course_enrollment_id"`
	EvaluationTypeID   string          `json:"evaluation_type_id"`
	EvaluationDate     sql.NullTime    `json:"evaluation_date"`
	FinalScore         sql.NullFloat64 `json:"final_score"`
}
