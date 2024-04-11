package userUsecase

import (
	"Ecom/pkg/repository/user/interfaces"
	services "Ecom/pkg/usecase/user/interfaces"
	"Ecom/pkg/utils/models"
	"errors"
)

type UProductUsecase struct {
	repo interfaces.UProductRepo
}

func NewUProductUsecase(repo interfaces.UProductRepo) services.UProductUsecase {
	return &UProductUsecase{
		repo: repo,
	}
}

func (Upro *UProductUsecase) ViewProduct(count int, page int, sort string) ([]models.ProductAdd, error) {
	products, err := Upro.repo.ViewProduct(count, page, sort)
	if err != nil {
		return []models.ProductAdd{}, err
	}

	return products, nil
}

func (Upro *UProductUsecase) ProductByName(Proname string) ([]models.ProductAdd, error) {
	exist, err := Upro.repo.CheckProductExistOrNot(Proname)
	if exist == 0 {
		return []models.ProductAdd{}, errors.New("product does not exists")
	}

	if err != nil {
		return []models.ProductAdd{}, err

	}

	product, err := Upro.repo.ProductByName(Proname)

	if err != nil {
		return []models.ProductAdd{}, err

	}

	return product, nil

}
