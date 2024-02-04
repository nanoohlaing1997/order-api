# order-api

Order API backend with golang

## Requirement

- go 1.21.0
- mysql 5.7

## Setup guide

1. Clone the repository
2. Copy `.env.example` to `.env`
3. Modify the configuration in `.env` as you required
4. Run `docker-compose up app`
5. The application will be server on port `8080`

## Environment config

`REST_PORT` : Port to serve the application

`ORDER_DB` : Database config

`GOOGLE_API_KEY` : Configuration of google map API key

## API Documentation

### Place order

- **Method**: `POST`
- **URL**: `/orders`
- **Feature**: - Calls google map API to measure the distance between origin and destination in meters and stores it in order table
- **Additional Info**: - if you encounter {"error": "Distance information not found"}. You must provide a valid API key. If you don't have valid API key to test you can remove the commented code block in `service\google_map_api.go` lines 29 to 41 and comment out lines 44 to 46. The commented code block provides dummy response from google map API.

### Take order

- **Method**: `PATCH`
- **URL**: `/orders/:id`
- **Feature**: - Retrieves the order by ID. Validate the order status is `UNASSIGNED` or `TAKEN`. If it is not `TAKEN`, the order is placed.
- **Additional Info**: - To avoid concurrency issues during update. Transactions are used to lock the table and if an error occurs during processing, transaction is rolled back.

### List order

- **Method**: `GET`
- **URL**: `/orders?page=:page&limit=:limit`
- **Feature**: - Retrieves list of data base on page and limit values.
