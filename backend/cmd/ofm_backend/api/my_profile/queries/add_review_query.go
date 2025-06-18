package queries

const AddReviewQuery = `
WITH allowed AS (
    SELECT 1 FROM orders O
    WHERE (O.order_id = $1 AND O.customer_id = $2)
)

INSERT INTO reviews (content, rating)
SELECT $3, $4
FROM allowed
RETURNING review_id
`