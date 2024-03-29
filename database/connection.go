package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseManager struct {
	*OrderDB
}

func DBConn(conn string) *gorm.DB {
	conn = conn + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewDatabaseManager(dbName string) *DatabaseManager {
	return &DatabaseManager{
		OrderManager(DBConn(dbName)),
	}
}
