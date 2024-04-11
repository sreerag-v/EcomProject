package adminUsecase

import (
	"Ecom/pkg/repository/admin/interfaces"
	services "Ecom/pkg/usecase/admin/interfaces"
	"Ecom/pkg/utils/models"
	"errors"
)

type ProductUsecase struct {
	repo interfaces.ProductRepo
}

func NewProductUsecase(repo interfaces.ProductRepo) services.ProductUsecase {
	return &ProductUsecase{
		repo: repo,
	}
}

func (pro *ProductUsecase) AddProduct(body models.ProductAdd) error {
	exist, err := pro.repo.CheckProductExistOrNot(body.ProductName)

	if exist != 0 {
		return errors.New("product already exists")
	}

	if err != nil {
		return err
	}

	err = pro.repo.AddProduct(body)

	if err != nil {
		return err
	}

	ProId, err := pro.repo.GetProductId(body.ProductName)
	if err != nil {
		return err
	}

	err = pro.repo.AddToInventory(ProId, body.Stock)
	if err != nil {
		return err
	}

	return nil

}

func (pro *ProductUsecase) ViewProduct(count int, page int, sort string) ([]models.ProductAdd, error) {
	products, err := pro.repo.ViewProduct(count, page, sort)
	if err != nil {
		return []models.ProductAdd{}, err
	}

	return products, nil
}

func (pro *ProductUsecase) ProductByName(Proname string) ([]models.ProductAdd, error) {
	exist, err := pro.repo.CheckProductExistOrNot(Proname)
	if exist == 0 {
		return []models.ProductAdd{}, errors.New("product does not exists")
	}

	if err != nil {
		return []models.ProductAdd{}, err

	}

	product, err := pro.repo.ProductByName(Proname)

	if err != nil {
		return []models.ProductAdd{}, err

	}

	return product, nil

}

func (pro *ProductUsecase) EditProduct(Proid int, body models.UpdateProduct) error {
	exist, err := pro.repo.CheckProductById(Proid)

	if err != nil {
		return err
	}

	if exist == 0 {
		return errors.New("product does not exists")
	}

	UintStock := uint(body.Stock)

	err = pro.repo.UpdateStock(UintStock, Proid)
	if err != nil {
		return err
	}

	NewBody := models.PostUpdate{
		Product_Name: body.Product_Name,
		Price:        body.Price,
		Color:        body.Color,
	}

	err = pro.repo.EditProduct(Proid, NewBody)
	if err != nil {
		return err
	}

	return nil

}

func (pro *ProductUsecase) DeleteProduct(proid int) error {
	exist, err := pro.repo.CheckProductById(proid)

	if err != nil {
		return err
	}

	if exist == 0 {
		return errors.New("product does not exists")
	}

	err = pro.repo.DeleteProduct(proid)
	if err != nil {
		return err
	}

	return nil
}
