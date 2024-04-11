package domain

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required,email" validate:"email"`
	Password  string `json:"password" gorm:"validate:required"`
	Previlege string `json:"previlege" gorm:"previlege:2;default:'normal_admin';previlage IN ('super_admin','normal_admin')"`
}

type Product struct {
	gorm.Model
	ProductId   uint   `json:"product_id" gorm:"autoIncrement"  `
	ProductName string `json:"product_name" gorm:"not null"  `
	Price       uint   `json:"price" gorm:"not null"  `
	Image       string `json:"image" gorm:"not null"  `
	Color       string `json:"color" gorm:"not null"  `
}

type Inventory struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Product   Product `gorm:"ForeignKey:product_id"`
	ProductId uint
	Stock     uint `json:"stock"  `
}

type Order struct {
	OrderId     uint    `json:"order_id" gorm:"primaryKey"`
	Product     Product `gorm:"ForeignKey:product_id"`
	ProductId   uint
	User        User `gorm:"ForeignKey:id"`
	Id          int
	OrderStatus string `json:"order_status"`
	Quantity    int    `json:"quantity"`
	CreatedAt   time.Time
}
