package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/nanoohlaing1997/order-api/database"
)

var validate = validator.New()

type Controller struct {
	dbm *database.DatabaseManager
}

func NewControllerManager(dbm *database.DatabaseManager) *Controller {
	return &Controller{
		dbm: dbm,
	}
}
