package domain

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Project struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	OrganizationID string    `json:"organization_id"`
	Environments   []string  `json:"environments"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewOrganization(name, description string) *Organization {
	now := time.Now().UTC()
	return &Organization{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func NewProject(name, description, organizationID string, environments []string) *Project {
	now := time.Now().UTC()
	if environments == nil {
		environments = []string{"development", "staging", "production"}
	}
	return &Project{
		ID:             uuid.New().String(),
		Name:           name,
		Description:    description,
		OrganizationID: organizationID,
		Environments:   environments,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
