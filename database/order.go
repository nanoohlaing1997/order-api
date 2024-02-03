package database

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint64    `json:"id"`
	Distance  float64   `json:"distance"`
	Status    string    `json:"status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
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
