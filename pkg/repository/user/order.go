package userRepo

import (
	"Ecom/pkg/domain"
	"Ecom/pkg/repository/user/interfaces"
	"Ecom/pkg/utils/models"

	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) interfaces.OrderRepo {
	return &OrderRepo{
		DB: db,
	}
}

func (odr *OrderRepo) CheckProductById(proid uint) (int64, error) {
	var count int64
	err := odr.DB.Raw("SELECT COUNT(*) FROM products WHERE id = $1", proid).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ord *OrderRepo) CheckQuantiyOfProduct(proid uint) (int64, error) {
	var count int64

	err := ord.DB.Table("inventories").Where("product_id = ?", proid).Select("stock").Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil

}

func (ord *OrderRepo) PlaceOrder(body models.PlaceOrderData, Uid int) error {
	tx := ord.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Create triggers if they don't exist
	if err := ord.createTriggers(); err != nil {
		tx.Rollback()
		return err
	}

	order := domain.Order{
		ProductId:   body.ProductId,
		Id:          Uid,
		Quantity:    body.Quantity,
		OrderStatus: "Order Placed",
	}

	// Place the order
	err := ord.DB.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// This Area Is Will Make A Little Bit Confunction -> Just Follow The Comments

// <...............................................................................>
// Function to create triggers if not exists
func (ord *OrderRepo) createTriggers() error {
	// Check if the trigger function exists
	var count int
	err := ord.DB.Raw("SELECT COUNT(*) FROM pg_proc WHERE proname = 'decrease_inventory_stock'").Scan(&count).Error
	if err != nil {
		return err
	}

	// If the trigger function exists, drop and recreate it
	if count > 0 {
		if err := ord.dropTriggerFunction(); err != nil {
			return err
		}
	}

	// Check if the trigger exists
	err = ord.DB.Raw("SELECT COUNT(*) FROM pg_trigger WHERE tgname = 'trigger_decrease_inventory_stock'").Scan(&count).Error
	if err != nil {
		return err
	}

	// If the trigger exists, drop and recreate it
	if count > 0 {
		if err := ord.dropTrigger(); err != nil {
			return err
		}
	}

	// Create trigger function
	if err := ord.createTriggerFunction(); err != nil {
		return err
	}

	// Create trigger
	if err := ord.createTrigger(); err != nil {
		return err
	}

	return nil
}

// Function to create trigger function if not exists
func (ord *OrderRepo) createTriggerFunction() error {
	err := ord.DB.Exec(`
        CREATE OR REPLACE FUNCTION decrease_inventory_stock()
        RETURNS TRIGGER AS
        $$
        BEGIN
            UPDATE inventories
            SET stock = stock - NEW.quantity
            WHERE product_id = NEW.product_id;

            RETURN NEW;
        END;
        $$
        LANGUAGE plpgsql;
    `).Error
	return err
}

// Function to create trigger if not exists
func (ord *OrderRepo) createTrigger() error {
	err := ord.DB.Exec(`
        CREATE TRIGGER trigger_decrease_inventory_stock
        AFTER INSERT ON orders
        FOR EACH ROW
        EXECUTE FUNCTION decrease_inventory_stock();
    `).Error
	return err
}

// Function to drop trigger function
func (ord *OrderRepo) dropTriggerFunction() error {
	// Drop trigger first
	if err := ord.dropTrigger(); err != nil {
		return err
	}

	// Now drop the function
	err := ord.DB.Exec(`
        DROP FUNCTION decrease_inventory_stock();
    `).Error
	return err
}

// Function to drop trigger
func (ord *OrderRepo) dropTrigger() error {
	err := ord.DB.Exec(`
        DROP TRIGGER trigger_decrease_inventory_stock ON orders;
    `).Error
	return err
}

//<...............................................................................>

func (ord *OrderRepo) ViewOrders(Uid, count, pageN int) ([]models.ViewOrders, error) {
	limit := count
	offset := (pageN - 1) * limit

	var orders []models.ViewOrders
	err := ord.DB.Table("orders").
		Select("orders.order_id,orders.product_id,products.product_name,products.price,products.color,products.image").
		Joins("LEFT JOIN products ON products.id = orders.product_id").
		Limit(limit).
		Offset(offset).
		Scan(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}
