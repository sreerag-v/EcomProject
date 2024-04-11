package domain

import "time"

type User struct {
	ID              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"required,number"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirmpassword" validate:"required,eqfield=Password"`
	CreatedAt       time.Time
}
