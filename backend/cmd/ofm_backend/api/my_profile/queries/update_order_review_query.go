package queries

const UpdateOrderReviewQuery = `
UPDATE orders O
SET review_id = $1
WHERE O.customer_id = $2 AND O.order_id = $3
`