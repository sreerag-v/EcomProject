package interfaces

import (
	"Ecom/pkg/utils/models"
)

type ProductUsecase interface {
	AddProduct(models.ProductAdd) error
	ViewProduct(count int, page int, sort string) ([]models.ProductAdd, error)
	ProductByName(Proname string) ([]models.ProductAdd, error)
	EditProduct(Proid int, body models.UpdateProduct) error
	DeleteProduct(proid int) error
}
