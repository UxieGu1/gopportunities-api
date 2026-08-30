package schemas

import (
	"time"

	"gorm.io/gorm"
)


type Application struct {
	gorm.Model
	CandidateID uint `gorm:"not null"`	
	Candidate   Candidate

	OpeningID 	uint `gorm:"not null"`
	Opening 	Opening

	Status      string    `gorm:"default:'APPLIED'"` // APPLIED, REVIEWING, INTERVIEW, REJECTED, HIRED
	Notes 		string
}

type ApplicationResponse struct {
	ID          uint      `json:"id"`
	CandidateID uint      `json:"candidateId"`
	OpeningID   uint      `json:"openingId"`
	Status      string    `json:"status"`
	Notes       string    `json:"notes"`
	CreatedAt 	time.Time `json:"createdAt"`
	UpdatedAt 	time.Time `json:"updatedAt"`
}