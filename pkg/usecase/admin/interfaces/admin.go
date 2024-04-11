package interfaces

import (
	"Ecom/pkg/domain"
	"Ecom/pkg/utils/models"
	"context"
	"time"
)

type AdminUsecase interface {
	AdminSignup(domain.Admin) error
	AdminLogin(context.Context, models.AdminLogin) (string, error)
	ViewAllOrders(count, page int) ([]models.AdminViewOrder, error)
	FetchOrderDates(fromtime, totime time.Time) ([]models.SalesReport, error)
}
