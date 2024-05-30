package models

import "time"

type Users struct {
	ID        uint64     `json:"id" gorm:"primary_key;auto_increment"`
	Username  string     `json:"username" binding:"required,min=3" gorm:"type:varchar(100);not null"`
	Email     string     `json:"email" binding:"required,email" gorm:"type:varchar(100);not null"`
	Password  string     `json:"password" binding:"required,min=6" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	Products  []Products `json:"products,omitempty" gorm:"foreignKey:UserID;references:ID"`
}
