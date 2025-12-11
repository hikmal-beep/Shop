package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64 	`gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string	`gorm:"type:varchar(100)" json:"name"`
	Email     string	`gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string	`gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Shop      *Shop 	`gorm:"foreignKey:UserID" json:"shop,omitempty"`
}

func (user *User) SetPassword(password string) error{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}