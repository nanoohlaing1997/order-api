package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nanoohlaing1997/order-api/api"
	"github.com/nanoohlaing1997/order-api/database"
	"github.com/stretchr/testify/assert"
)

func setup() (*api.Controller, *database.DatabaseManager) {
	// Use hard code for order_db_test db connection
	dbConnTest := "root:root@tcp(mysql)/order_db_test"
	testDBM := database.NewDatabaseManager(dbConnTest)
	controller := api.NewControllerManager(testDBM)

	return controller, testDBM
}

func TestCreateOrder404(t *testing.T) {
	controller, _ := setup()

	reqBody := api.CreateOrderRequest{
		Origin:      []string{"0.0", "0.0"},
		Destination: []string{"1.0", "1.0"},
	}

	requestBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/orders", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.CreateOrder)
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Result().StatusCode, http.StatusNotFound)
}

func TestCreateOrder400(t *testing.T) {
	controller, _ := setup()

	reqBody := map[string][]string{
		"origin": {"0.0", "0.0"},
		"dest":   {"1.1", "1.0"},
	}

	requestBody, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/orders", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.CreateOrder)
	handler.ServeHTTP(res, req)

	response := api.CreateOrderResponse{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Result().StatusCode, http.StatusBadRequest)
}

func TestTakeOrder200(t *testing.T) {
	controller, dbm := setup()
	order, _ := dbm.CreateOrder(&database.Order{
		Distance: 10000,
		Status:   api.Unassign,
	})

	defer dbm.Truncate()

	orderID := strconv.FormatUint(order.ID, 10)
	reqBody := api.TakeOrderRequestAndResponse{
		Status: api.Taken,
	}
	requestBody, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(
		"PATCH",
		fmt.Sprintf("/orders/%s", orderID),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": orderID})
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.TakeOrder)
	handler.ServeHTTP(res, req)

	response := api.TakeOrderRequestAndResponse{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.Result().StatusCode, http.StatusOK)
	assert.Equal(t, response.Status, "SUCCESS")
}

func TestTakeOrder404(t *testing.T) {
	controller, dbm := setup()
	order, _ := dbm.CreateOrder(&database.Order{
		Distance: 5000,
		Status:   api.Taken,
	})

	defer dbm.Truncate()

	orderID := strconv.FormatUint(order.ID, 10)
	reqBody := api.TakeOrderRequestAndResponse{
		Status: api.Taken,
	}
	requestBody, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(
		"PATCH",
		fmt.Sprintf("/orders/%s", orderID),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": orderID})
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.TakeOrder)
	handler.ServeHTTP(res, req)

	assert.Equal(t, res.Result().StatusCode, http.StatusNotFound)
}

func TestListOrder200(t *testing.T) {
	controller, dbm := setup()
	_, _ = dbm.CreateOrder(&database.Order{
		Distance: 5000,
		Status:   api.Unassign,
	})

	_, _ = dbm.CreateOrder(&database.Order{
		Distance: 5000,
		Status:   api.Taken,
	})

	defer dbm.Truncate()

	req, err := http.NewRequest("GET", "/orders?page=1&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"page": "1", "limit": "10"})
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.ListOrder)
	handler.ServeHTTP(res, req)
	fmt.Println(res.Body)
	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response []*database.Order
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(response), 2)
}

func TestListOrderEmpty(t *testing.T) {
	controller, _ := setup()
	req, err := http.NewRequest("GET", "/orders?page=1&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"page": "1", "limit": "10"})
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.ListOrder)
	handler.ServeHTTP(res, req)
	fmt.Println(res.Body)
	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response []*database.Order
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(response), 0)
}
