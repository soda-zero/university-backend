package models

type Career struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	DurationYears int    `json:"duration_years"`
	DepartmentID  string `json:"department_id"`
	CareerLevelID string `json:"career_level_id"`
}
