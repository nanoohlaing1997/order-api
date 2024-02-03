package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nanoohlaing1997/order-api/database"
)

type CreateOrderRequest struct {
	Origin      []string `json:"origin"`
	Destination []string `json:"destination"`
}

type CreateOrderResponse struct {
	ID       uint64  `json:"id"`
	Distance float64 `json:"distance"`
	Status   string  `json:"status"`
}

type TakeOrderRequestAndResponse struct {
	Status string `json:"status"`
}

type ListOrderResponse struct {
	Orders []*database.Order
}

type ErrorResponse struct {
	Message string `json:"error"`
}

func isValidLatLong(latString, longString string) bool {
	lat, err := strconv.ParseFloat(latString, 64)
	if err != nil {
		return false
	}

	long, err := strconv.ParseFloat(longString, 64)
	if err != nil {
		return false
	}

	return (lat >= -90 && lat <= 90) && (long >= -180 && long <= 180)
}

func returnError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Message: message,
	})
}
