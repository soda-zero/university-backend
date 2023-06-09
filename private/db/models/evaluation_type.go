package models

type EvaluationType struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	IsGroupActivity bool   `json:"is_group_activity"`
}
