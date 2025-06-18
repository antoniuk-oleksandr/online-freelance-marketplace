package queries

const AddOrderDeliveryQuery = `
WITH allowed AS (
  SELECT 1
  FROM orders
  WHERE order_id = $1 AND freelancer_id = $3
)
INSERT INTO orders_deliveries (order_id, message)
SELECT $1, $2
FROM allowed
RETURNING delivery_id
`
