package interfaces

import "Ecom/pkg/utils/models"

type OrderUsecase interface {
	PlaceOrder(body models.PlaceOrderData, Uid int) error
	ViewOrders(Uid, count, pageN int) ([]models.ViewOrders, error)
}
