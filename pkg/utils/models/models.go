package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AdminLogin struct {
	Email    string `json:"email" gorm:"validate:required,email" validate:"email"`
	Password string `json:"password"`
}

type AuthCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type AdminDetailsResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name" `
	Email     string `json:"email" `
	Previlege string `json:"previlege"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostUpdate struct {
	Product_Name string `json:"product_name"`
	Price        int    `json:"price"`
	Color        string `json:"color"`
}

type ProductAdd struct {
	ProductName string `json:"product_name" gorm:"not null"  `
	Price       uint   `json:"price" gorm:"not null"  `
	Image       string `json:"image" gorm:"not null"  `
	Stock       uint   `json:"stock"  `
	Color       string `json:"color" gorm:"not null"  `
}

type UpdateProduct struct {
	Product_Name string `json:"product_name"`
	Price        int    `json:"price"`
	Color        string `json:"color"`
	Stock        int    `json:"stock"`
}

type PlaceOrderData struct {
	ProductId uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type ViewOrders struct {
	Order_id     uint
	Product_id   int
	Product_name string
	Price        int
	Color        string
	Image        string
}

type AdminViewOrder struct {
	Order_id     uint
	User_id      int
	User_name    string
	Product_id   int
	Product_name string
	Price        int
	Color        string
	Image        string
}

type SalesReport struct {
	CreatedAt   time.Time
	OrderID     uint
	ProductName string
	Price       uint
	UserName    string
	OrderStatus string
	Quantity    int
}
