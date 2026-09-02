package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	UserID      uint   `gorm:"index"`
	Name        string `gorm:"not null"`
	Description string
	Website     string
	Email       string
	Location    string
	Openings    []Opening
}

type CompanyResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	Email       string    `json:"email"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
