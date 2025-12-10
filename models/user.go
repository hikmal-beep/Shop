package models

import "time"

type User struct {
	ID        int64 	`gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string	`gorm:"type:varchar(100)" json:"name"`
	Email     string	`gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string	`gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Shops     []Shop 	`gorm:"foreignKey:UserID" json:"tokos,omitempty"`
}