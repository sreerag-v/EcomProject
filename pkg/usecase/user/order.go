package userUsecase

import (
	"Ecom/pkg/repository/user/interfaces"
	services "Ecom/pkg/usecase/user/interfaces"
	"Ecom/pkg/utils/models"
	"errors"
)

type OrderUsecase struct {
	repo interfaces.OrderRepo
}

func NewOrderUsecase(repo interfaces.OrderRepo) services.OrderUsecase {
	return &OrderUsecase{
		repo: repo,
	}
}

func (ord *OrderUsecase) PlaceOrder(body models.PlaceOrderData, Uid int) error {
	exist, err := ord.repo.CheckProductById(body.ProductId)
	if err != nil {
		return err
	}

	if exist == 0 {
		return errors.New("product does not exists")
	}

	checkQu, err := ord.repo.CheckQuantiyOfProduct(body.ProductId)
	if body.Quantity > int(checkQu) || checkQu == 0 {
		return errors.New("product out of stock")
	}

	if err != nil {
		return err
	}

	err = ord.repo.PlaceOrder(body, Uid)
	if err != nil {
		return err
	}

	return nil
}

func (ord *OrderUsecase) ViewOrders(Uid, count, pageN int) ([]models.ViewOrders, error) {
	orders, err := ord.repo.ViewOrders(Uid, count, pageN)
	if err != nil {
		return []models.ViewOrders{}, nil
	}

	return orders, nil
}
