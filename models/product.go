package models

import "time"

type Product struct {
	ID          int64		`gorm:"primaryKey;autoIncrement" json:"id"`
	ShopID      int64		`gorm:"not null;index" json:"toko_id"`
	Product     string		`gorm:"type:varchar(255);not null" json:"product"`
	Description string		`gorm:"type:text" json:"description"`
	Quantity    int			`gorm:"not null" json:"quantity"`
	CreatedAt   time.Time	`json:"created_at" `
	UpdatedAt	time.Time	`json:"updated_at"`
	Shop		*Shop		`gorm:"foreignKey:ShopID" json:"toko,omitempty"`
}