package models

import "time"

type Shop struct {
	ID        	int64		`gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    	int64		`gorm:"not null;index"  json:"user_id"`
	Name      	string		`gorm:"type:varchar(100);not null" json:"name"`
	Address   	string		`gorm:"type:text" json:"address" `
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	Products 	[]Product	`gorm:"foreignKey:ShopID" json:"product,omitempty"`
}