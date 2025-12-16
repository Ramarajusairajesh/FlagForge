package domain

import (
	"time"

	"github.com/google/uuid"
)

type FlagType string

const (
	BooleanFlag FlagType = "boolean"
	StringFlag  FlagType = "string"
	NumberFlag  FlagType = "number"
	JSONFlag    FlagType = "json"
)

type FlagStatus string

const (
	FlagStatusActive   FlagStatus = "active"
	FlagStatusArchived FlagStatus = "archived"
)

type Flag struct {
	ID          string    `json:"id"`
	Name        string    `json:
ame"`
	Key         string    `json:"key"`
	Description string    `json:"description"`
	Type        FlagType  `json:"type"`
	Status      string    `json:"status"`
	DefaultOn   bool      `json:"default_on"`
	DefaultOff  bool      `json:"default_off"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	ProjectID   string    `json:"project_id"`
	Environment string    `json:"environment"`
}

type TargetingRule struct {
	ID          string                 `json:"id"`
	FlagID      string                 `json:"flag_id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Conditions  []Condition            `json:"conditions"`
	Variations  map[string]interface{} `json:"variations"`
	Rollout     Rollout                `json:"rollout"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

type Condition struct {
	Attribute string      `json:"attribute"`
	Operator  string      `json:"operator"` // equals, not_equals, contains, etc.
	Value     interface{} `json:"value"`
}

type Rollout struct {
	Percentage float64 `json:"percentage"`
	Seed       int64   `json:"seed"`
}

func NewFlag(name, key, description string, flagType FlagType, createdBy, projectID, environment string) *Flag {
	now := time.Now().UTC()
	return &Flag{
		ID:          uuid.New().String(),
		Name:        name,
		Key:         key,
		Description: description,
		Type:        flagType,
		Status:      string(FlagStatusActive),
		DefaultOn:   false,
		DefaultOff:  true,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
		ProjectID:   projectID,
		Environment: environment,
	}
}
