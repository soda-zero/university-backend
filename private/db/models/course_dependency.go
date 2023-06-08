package models

type CourseDependency struct {
	CourseID         string `json:"course_id"`
	RequiredCourseID string `json:"required_course_id"`
}
