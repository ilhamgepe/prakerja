package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Products struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" binding:"required,min=3" gorm:"type:varchar(100);not null"`
	Price     int32     `json:"price" binding:"required,min=1000" gorm:"not null;type:int"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID    uint64    `json:"user_id" gorm:"column:user_id;not null"`
}

func (p *Products) Validate() error {
	// err := validate.Struct(p).Error()
	return nil
}
