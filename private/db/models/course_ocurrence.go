package models

import "time"

type CourseOcurrence struct {
	ID            string    `json:"id"`
	OcurrenceYear int       `json:"ocurrence_year"`
	Code          string    `json:"code"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Capacity      int       `json:"capacity"`
	CourseID      string    `json:"course_id"`
	ProfessorID   string    `json:"professor_id"`
}
