package interfaces

import "Ecom/pkg/utils/models"

type OrderRepo interface {
	CheckProductById(Proid uint) (int64, error)
	CheckQuantiyOfProduct(proid uint) (int64, error)
	PlaceOrder(body models.PlaceOrderData, Uid int) error
	ViewOrders(Uid, count, pageN int) ([]models.ViewOrders, error)
}
