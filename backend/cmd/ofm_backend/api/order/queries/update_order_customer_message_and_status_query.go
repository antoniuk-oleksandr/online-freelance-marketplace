package queries

const AddOrderCustomerMessageAndStatusQuery = `
UPDATE orders
SET customer_message = $1, status_id = $2
WHERE order_id = $3
`
