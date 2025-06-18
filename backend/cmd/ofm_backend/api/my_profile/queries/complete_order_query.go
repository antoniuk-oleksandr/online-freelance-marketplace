package queries

const CompleteOrderQuery = `
UPDATE orders
SET
  ended_at = CURRENT_TIMESTAMP,
  status_id = 3
WHERE order_id = $1
AND freelancer_id = $2
`
