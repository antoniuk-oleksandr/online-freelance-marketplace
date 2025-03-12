package queries

const CheckIfOrderRequirementsSubmittedQuery = `
SELECT status_id = $1 AS submitted
FROM orders
WHERE order_id = $2
`