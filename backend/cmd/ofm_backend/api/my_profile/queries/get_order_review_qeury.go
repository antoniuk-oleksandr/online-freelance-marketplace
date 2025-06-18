package queries

const GetOrderReviewQuery = `
SELECT
    R.content,
    R.rating,
    json_build_object(
        'id', U.user_id,
        'username', U.username,
        'firstName', U.first_name,
        'surname', U.surname,
        'avatar', COALESCE(F.name, '')
    )AS customer
FROM orders O
JOIN reviews R ON R.review_id = O.review_id
LEFT JOIN users U ON U.user_id = O.customer_id
LEFT JOIN files F ON F.file_id = U.avatar_id
WHERE O.order_id = $1
`