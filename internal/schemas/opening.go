package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role      string
	CompanyID uint `gorm:"not null"`
	Company   Company
	Status    string `gorm:"not null;default:'OPEN'"`
	Location  string
	Remote    bool
	Link      string
	Salary    int64
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	CompanyID uint      `json:"companyId"`
	Status    string    `json:"status"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}
