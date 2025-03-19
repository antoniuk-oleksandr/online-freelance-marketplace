package queries

const GetMyProfileOrdersQuery = `
WITH orders_data AS (
    SELECT O.order_id, S.title, O.status_id as status, P.price,
    TO_CHAR(O.created_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"') AS date,
        (
            SELECT F.name
            FROM files F
            LEFT JOIN services_files SF ON SF.file_id = F.file_id AND SF.service_id = S.service_id
            ORDER BY F.file_id
            LIMIT 1
        ) as image
    FROM orders O
    LEFT JOIN services S ON s.service_id = O.service_id
    LEFT JOIN packages P ON p.package_id = O.service_package_id
    WHERE O.customer_id = $1
    ORDER BY O.created_at DESC , O.order_id
    OFFSET $2
    LIMIT $3
),
total_pages_data AS (
    SELECT CEIL(CAST(count(O.order_id) AS FLOAT) / $3) as total_pages
    FROM orders O
    WHERE customer_id = $1
)

SELECT
COALESCE(
    (SELECT json_agg(json_build_object(
        'order_id', orders_data.order_id,
        'title', orders_data.title,
        'status', orders_data.status,
        'price', orders_data.price,
        'date', orders_data.date,
        'image', orders_data.image
    )) FROM orders_data),
    '[]'::json
) as orders_data,
(
    SELECT * FROM total_pages_data
)
`