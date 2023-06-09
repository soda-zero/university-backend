package models

import "time"

type Schedule struct {
	ID                string    `json:"id"`
	CourseOcurrenceID string    `json:"course_ocurrence_id"`
	DayOfWeek         string    `json:"day_of_week"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	RoomID            string    `json:"room_id"`
}
