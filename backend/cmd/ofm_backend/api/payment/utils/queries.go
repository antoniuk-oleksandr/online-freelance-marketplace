package utils

const PaymentQuery = `
INSERT INTO payments (card_number, status_id)
VALUES ($1, $2)
RETURNING payment_id
`

const UpdatePaymentStatusQuery = `
UPDATE payments SET status_id = $1
WHERE payment_id = $2
`

const AddOrderQuery = `
INSERT INTO orders(
    customer_id, freelancer_id, service_id,
    service_package_id, status_id, payment_id
)
SELECT u.user_id, s.freelancer_id, s.service_id,
       $3, $4, $5
FROM users u, services s
WHERE u.username = $1 AND s.service_id = $2
RETURNING order_id
`

const GetUserEmailByUsernameQuery = `
SELECT email
FROM users
WHERE username = $1
`

const GetPaymentReceiptQuery = `
SELECT
    O.order_id,
    S.title AS service_name,
    P.title AS package_name,
    P.price,
    round((P.price * $1)::numeric, 2) AS service_fee,
    round((P.price + P.price * $1):: numeric, 2) AS total,
    O.created_at as date,
    U.email as customer_email
FROM orders O
INNER JOIN users U ON u.user_id = O.customer_id
INNER JOIN services S ON S.service_id = O.service_id
INNER JOIN packages P ON P.package_id = O.service_package_id
WHERE order_id = $2;
`
