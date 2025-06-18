package queries

const UpdateOrderStatusByOrderIdQuery = `
UPDATE orders
SET status_id = $1
WHERE order_id = $2 AND freelancer_id = $3
`
