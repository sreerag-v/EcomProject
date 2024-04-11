package interfaces

import "Ecom/pkg/utils/models"

type UProductUsecase interface {
	ViewProduct(count int, page int, sort string) ([]models.ProductAdd, error)
	ProductByName(Proname string) ([]models.ProductAdd, error)
}
