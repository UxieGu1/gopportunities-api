package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	UserID    uint   `gorm:"uniqueIndex;not null"` 
	User      User   
	Name      string `gorm:"not null"`
	Linkedin  string
	ResumeURL string 
	Skills    string 
}

type CandidateResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"userId"`
	Name      string `json:"name"`
	Linkedin  string `json:"linkedin"`
	ResumeURL string `json:"resumeUrl"`
	Skills    string `json:"skills"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}