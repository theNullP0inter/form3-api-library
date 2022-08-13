package models

import (
	"github.com/google/uuid"
)

type Account struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             uuid.UUID          `json:"id,omitempty"`
	OrganisationID uuid.UUID          `json:"organisation_id,omitempty"`
	Type           string             `json:"type" validate:"required"`
	Version        int64              `json:"version,omitempty"`
}
