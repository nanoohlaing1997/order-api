package database

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint64    `json:"id"`
	Distance  float64   `json:"distance"`
	Status    string    `json:"status"`
	CreatedAt time.Time `                gorm:"column:created_at"`
	UpdatedAt time.Time `                gorm:"column:updated_at"`
}

type OrderDB struct {
	db    *gorm.DB
	model *Order
}

func (order *Order) TableName() string {
	return "orders"
}

func OrderManager(db *gorm.DB) *OrderDB {
	return &OrderDB{
		db:    db,
		model: &Order{},
	}
}

func (odb *OrderDB) CreateOrder(order *Order) (*Order, error) {
	if res := odb.db.Create(&order); res != nil && res.RowsAffected <= 0 {
		return nil, res.Error
	}
	return order, nil
}

func (odb *OrderDB) TakeOrder(orderID uint64, status string) error {
	var order *Order
	tx := odb.db.Begin()

	if res := tx.Where("id = ?", orderID).First(&order); res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if order.Status == status {
		tx.Rollback()
		return errors.New("Order is already taken")
	}

	order.Status = status
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
