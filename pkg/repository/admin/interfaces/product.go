package interfaces

import (
	"Ecom/pkg/utils/models"
)

type ProductRepo interface {
	CheckProductExistOrNot(name string) (int64, error)
	AddProduct(body models.ProductAdd) error
	ViewProduct(count int, page int, sort string) ([]models.ProductAdd, error)
	ProductByName(Proname string) ([]models.ProductAdd, error)
	CheckProductById(proid int) (int64, error)
	EditProduct(proid int, body models.PostUpdate) error
	DeleteProduct(proid int) error

	AddToInventory(proid uint, stock uint) error
	GetProductId(name string) (uint, error)

	UpdateStock(stock uint, proid int) error
}
