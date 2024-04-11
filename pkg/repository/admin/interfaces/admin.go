package interfaces

import (
	"Ecom/pkg/domain"
	"Ecom/pkg/utils/models"
	"context"
	"time"
)

type AdminRepo interface {
	FindAdminByEmail(Email string) (int64, error)
	AdminSignup(body domain.Admin) error
	GetAdminDetailsByEmail(ctx context.Context, email string) (domain.Admin, error)
	ViewAllOrders(count, page int) ([]models.AdminViewOrder, error)
	FetchOrderDates(fromtime, totime time.Time) ([]models.SalesReport, error)
}
