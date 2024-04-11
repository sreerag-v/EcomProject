package adminRepo

import (
	"Ecom/pkg/domain"
	"Ecom/pkg/repository/admin/interfaces"
	"Ecom/pkg/utils/models"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) interfaces.AdminRepo {
	return &AdminRepo{
		DB: db,
	}
}

func (adm *AdminRepo) FindAdminByEmail(email string) (int64, error) {
	var count int64
	err := adm.DB.Raw("SELECT COUNT(*) FROM admins WHERE email = $1", email).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (adm *AdminRepo) AdminSignup(body domain.Admin) error {
	tx := adm.DB.Begin()
	//insert new admin to admin table
	err := tx.Exec("INSERT INTO admins (name,email,password) VALUES($1,$2,$3)", body.Name, body.Email, body.Password).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (adm *AdminRepo) GetAdminDetailsByEmail(ctx context.Context, email string) (domain.Admin, error) {
	var model domain.Admin
	err := adm.DB.Table("admins").Where("email = ?", email).Scan(&model).Error
	if err != nil {
		return domain.Admin{}, err
	}
	if ctx.Err() != nil {
		return domain.Admin{}, errors.New("timeout")
	}

	return model, nil
}

func (adm *AdminRepo) ViewAllOrders(count, page int) ([]models.AdminViewOrder, error) {
	limit := count
	offset := (page - 1) * limit

	var orders []models.AdminViewOrder
	err := adm.DB.Table("orders").
		Select("orders.order_id,users.id AS user_id,users.name AS user_name,orders.product_id,products.product_name,products.price,products.color,products.image").
		Joins("LEFT JOIN users ON users.id = orders.id").
		Joins("LEFT JOIN products ON products.id = orders.product_id").
		Limit(limit).
		Offset(offset).
		Scan(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (adm *AdminRepo) FetchOrderDates(fromtime, totime time.Time) ([]models.SalesReport, error) {
	var SalesReport []models.SalesReport
	err := adm.DB.Table("orders").
		Select("orders.created_at, orders.order_id, products.product_name, products.price, users.name AS user_name, orders.order_status, orders.quantity").
		Joins("JOIN products ON orders.product_id = products.product_id").
		Joins("JOIN users ON orders.id = users.id").
		Where("orders.created_at BETWEEN ? AND ?", fromtime, totime).
		Scan(&SalesReport).Error

	if err != nil {
		return []models.SalesReport{}, err
	}

	return SalesReport, nil
}
