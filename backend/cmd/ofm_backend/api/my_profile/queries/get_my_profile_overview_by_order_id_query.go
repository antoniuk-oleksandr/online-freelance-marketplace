package queries

const GetMyProfileOverviewByOrderIdQuery = `
SELECT
    O.order_id,
    O.accepted_at + MAKE_INTERVAL(days := P.delivery_days) AS delivery_date,
    O.created_at,
    P.price AS subtotal,
    round((P.price * $2)::numeric, 2) AS service_fee,
    round((P.price + P.price * $2):: numeric, 2) AS total_price,
    O.status_id AS status,
    json_build_object(
        'image', (
            SELECT COALESCE(F.name, '')
            FROM services_files SF
            LEFT JOIN files F ON F.file_id = SF.file_id
            WHERE service_id = o.service_id
            ORDER BY SF.file_id
            LIMIT 1
        ),
        'title', S.title,
        'package', json_build_object(
            'description', P.description,
            'name', P.title,
            'delivery_time', P.delivery_days
        )
    ) as service,
    json_build_object(
        'username', U.username,
        'id', U.user_id,
        'avatar', COALESCE(F1.name, '')
    ) as freelancer
FROM orders O
LEFT JOIN packages P ON P.package_id = O.service_package_id
LEFT JOIN services S ON S.service_id = O.service_id
LEFT JOIN users U ON user_id = O.freelancer_id
LEFT JOIN files F1 ON F1.file_id = U.avatar_id
WHERE O.order_id = $1
AND (O.freelancer_id = $3 OR O.customer_id = $3)
`
