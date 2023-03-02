package models

import (
	"chain-vote-api/enums"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Election struct {
	gorm.Model
	ID          uuid.UUID            `gorm:"type:uuid;primary_key"`
	Name        string               `gorm:"size:255;not null;" json:"name"`
	Description string               `gorm:"size:255;not null;" json:"description"`
	StartDate   time.Time            `gorm:"size:255;not null;" json:"start_date"`
	EndDate     time.Time            `gorm:"size:255;not null;" json:"end_date"`
	Status      enums.ElectionStatus `gorm:"not null;" json:"status"`
	CreatedBy   User                 `gorm:"foreignkey:CreatedByID" json:"created_by"`
	CreatedByID uuid.UUID            `gorm:"type:uuid;not null" json:"created_by_id"`
}

type ElectionListing struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Status      string      `json:"status"`
	Creator     UserListing `json:"created_by"`
}

type CreateElectionInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`

}

func (e *Election) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}
