package models

type Course struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Year       int    `json:"year"`
	Semester   int    `json:"semester"`
	Optative   bool   `json:"optative"`
	CourseCode string `json:"course_code"`
	CareerID   string `json:"career_id"`
}
