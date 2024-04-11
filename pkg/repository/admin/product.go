package adminRepo

import (
	"Ecom/pkg/domain"
	"Ecom/pkg/repository/admin/interfaces"
	"Ecom/pkg/utils/models"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) interfaces.ProductRepo {
	return &ProductRepo{
		DB: db,
	}
}

func (pro *ProductRepo) CheckProductExistOrNot(name string) (int64, error) {
	var count int64
	err := pro.DB.Raw("SELECT COUNT(*) FROM products WHERE product_name = $1", name).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (pro *ProductRepo) AddProduct(body models.ProductAdd) error {
	tx := pro.DB.Begin()
	Pr := domain.Product{
		ProductName: body.ProductName,
		Price:       body.Price,
		Image:       body.Image,
		Color:       body.Color,
	}
	//create product table
	err := pro.DB.Create(&Pr).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (pro *ProductRepo) AddToInventory(proid uint, stock uint) error {
	err := pro.DB.Exec("INSERT INTO inventories(product_id,stock)VALUES($1,$2)", proid, stock).Error
	if err != nil {
		return err
	}

	return nil
}

func (pro *ProductRepo) ViewProduct(count int, page int, sort string) ([]models.ProductAdd, error) {
	limit := count
	offset := (page - 1) * limit

	var products []models.ProductAdd
	err := pro.DB.Table("products").
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

func (pro *ProductRepo) GetProductId(name string) (uint, error) {
	var ProId uint
	err := pro.DB.Table("products").Where("product_name = ?", name).Select("id").Scan(&ProId).Error
	if err != nil {
		return 0, err
	}
	return ProId, nil
}

func (pro *ProductRepo) ProductByName(Proname string) ([]models.ProductAdd, error) {
	var Product []models.ProductAdd
	err := pro.DB.Table("products").Where("product_name = ?", Proname).
		Select("product_name,price,color,stock,image").
		Joins("LEFT JOIN inventories ON products.id = inventories.product_id").
		Where("products.deleted_at IS NULL").
		Scan(&Product).Error

	if err != nil {
		return []models.ProductAdd{}, err
	}
	return Product, nil

}

func (pro *ProductRepo) CheckProductById(proid int) (int64, error) {
	var count int64
	err := pro.DB.Raw("SELECT COUNT(*) FROM products WHERE id = $1", proid).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (pro *ProductRepo) DeleteProduct(proid int) error {
	tx := pro.DB.Begin()

	// Delete the product with the given ID
	err := tx.Delete(&domain.Product{}, proid).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (pro *ProductRepo) EditProduct(proid int, body models.PostUpdate) error {
	tx := pro.DB.Begin()
	err := pro.DB.Table("products").Where("id = ?", proid).Updates(body).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (pro *ProductRepo) UpdateStock(stock uint, proid int) error {
	tx := pro.DB.Begin()

	// Use Model method to specify the model to update and the condition
	err := pro.DB.Model(&domain.Inventory{}).
		Where("product_id = ?", proid).
		UpdateColumn("stock", stock).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
