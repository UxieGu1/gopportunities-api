package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Opening struct {
	gorm.Model
	Role     string
	CompanyID uint `gorm:"not null"`
	Company  Company
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	CompanyID uint    	`json:"companyId"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}
