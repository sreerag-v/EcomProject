package interfaces

import "Ecom/pkg/utils/models"

type UProductRepo interface {
	ViewProduct(count int, page int, sort string) ([]models.ProductAdd, error)
	CheckProductExistOrNot(name string) (int64, error)
	ProductByName(Proname string) ([]models.ProductAdd, error)
}
