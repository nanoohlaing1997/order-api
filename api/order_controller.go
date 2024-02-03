package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanoohlaing1997/order-api/database"
	"github.com/nanoohlaing1997/order-api/service"
)

const (
	Unassign = "UNASSIGNED"
	Taken    = "TAKEN"
)

func (c *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		returnError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use validator to perform struct validation
	if err := validate.Struct(reqBody); err != nil {
		returnError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request origin and destination have lat and long values
	if len(reqBody.Origin) != 2 || len(reqBody.Destination) != 2 {
		returnError(
			w,
			"Origin and destination must contain exactly 2 values",
			http.StatusBadRequest,
		)
		return
	}

	// Validate lat long values
	if (!isValidLatLong(reqBody.Origin[0], reqBody.Origin[1])) ||
		!isValidLatLong(reqBody.Destination[0], reqBody.Destination[1]) {
		returnError(
			w,
			"Origin and destination lat and long values are not valid",
			http.StatusNotAcceptable,
		)
		return
	}

	// Get distance between origin and destination using Google API
	distance, err := service.GetDistance(reqBody.Origin, reqBody.Destination)
	if err != nil {
		returnError(w, err.Error(), http.StatusNotFound)
		return
	}

	// Create order
	orderObj, err := c.dbm.CreateOrder(&database.Order{
		Distance: distance,
		Status:   Unassign,
	})
	if err != nil {
		returnError(w, err.Error(), http.StatusNotFound)
		return
	}

	res := &CreateOrderResponse{
		ID:       orderObj.ID,
		Distance: distance,
		Status:   orderObj.Status,
	}

	json.NewEncoder(w).Encode(res)
}

func (c *Controller) TakeOrder(w http.ResponseWriter, r *http.Request) {
	// Validate URL parameter is valid or not
	vars := mux.Vars(r)
	stringOrderID := vars["id"]
	orderID, _ := service.StringToUint64(stringOrderID)
	if orderID <= 0 {
		returnError(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// Validate the json body request
	var reqBody TakeOrderRequestAndResponse
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		returnError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the request format
	if err := validate.Struct(reqBody); err != nil {
		returnError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if reqBody.Status != Taken {
		returnError(w, "Status is invalid", http.StatusNotAcceptable)
		return
	}
	if err := c.dbm.TakeOrder(orderID, reqBody.Status); err != nil {
		returnError(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(TakeOrderRequestAndResponse{Status: "SUCCESS"})
}
