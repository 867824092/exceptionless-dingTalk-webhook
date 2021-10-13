package model

import "time"

// ExceptionlessWebHookModel webhook 实体
type ExceptionlessWebHookModel struct {
	Id               string    `json:"id"`
	Url              string    `json:"url"`
	OccurrenceDate   time.Time `json:"occurrence_date"`
	Type             string    `json:"type"`
	Message          string    `json:"message"`
	ProjectId        string    `json:"project_id"`
	ProjectName      string    `json:"project_name"`
	OrganizationId   string    `json:"organization_id"`
	OrganizationName string    `json:"organization_name"`
	StackId          string    `json:"stack_id"`
	StackUrl         string    `json:"stack_url"`
	StackTitle       string    `json:"stack_title"`
	StackTags        []string  `json:"stack_tags"`
	TotalOccurrences int       `json:"total_occurrences"`
	FirstOccurrence  time.Time `json:"first_occurrence"`
	LastOccurrence   time.Time `json:"last_occurrence"`
	IsNew            bool      `json:"is_new"`
	IsRegression     bool      `json:"is_regression"`
	IsCritical       bool      `json:"is_critical"`
}
