package queries

const GetOrderStatusByOrderIdQuery = `
SELECT status_id
FROM orders
WHERE order_id = $1
`