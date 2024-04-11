package userRepo

import (
	"Ecom/pkg/repository/user/interfaces"
	"Ecom/pkg/utils/models"

	"gorm.io/gorm"
)

type UProductRepo struct {
	DB *gorm.DB
}

func NewUProductRepo(db *gorm.DB) interfaces.UProductRepo {
	return &UProductRepo{
		DB: db,
	}
}

func (Upro *UProductRepo) ViewProduct(count int, page int, sort string) ([]models.ProductAdd, error) {
	limit := count
	offset := (page - 1) * limit

	var products []models.ProductAdd
	err := Upro.DB.Table("products").
		Select("products.product_name, products.price, products.color, products.image, inventories.stock").
		Joins("LEFT JOIN inventories ON products.id = inventories.product_id").
		Where("products.deleted_at IS NULL").
		Order("products.price " + sort).
		Limit(limit).
		Offset(offset).
		Scan(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (Upro *UProductRepo) CheckProductExistOrNot(name string) (int64, error) {
	var count int64
	err := Upro.DB.Raw("SELECT COUNT(*) FROM products WHERE product_name = $1", name).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (Upro *UProductRepo) ProductByName(Proname string) ([]models.ProductAdd, error) {
	var Product []models.ProductAdd
	err := Upro.DB.Table("products").Where("product_name = ?", Proname).
		Select("products.product_name, products.price, products.color, products.image, inventories.stock").
		Joins("LEFT JOIN inventories ON products.id = inventories.product_id").
		Where("products.deleted_at IS NULL").
		Scan(&Product).Error

	if err != nil {
		return []models.ProductAdd{}, err
	}
	return Product, nil
}
