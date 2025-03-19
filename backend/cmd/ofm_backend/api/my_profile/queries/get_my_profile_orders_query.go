package queries

const GetMyProfileOrdersQuery = `
WITH orders_data AS (
    SELECT
        O.order_id,
        S.title,
        O.status_id AS status,
        P.price,
        TO_CHAR(O.created_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"') AS date,
        F.name AS image
    FROM orders O
    LEFT JOIN services S ON S.service_id = O.service_id
    LEFT JOIN services_files SF ON SF.service_id = S.service_id
    LEFT JOIN files F ON F.file_id = SF.file_id
    LEFT JOIN packages P ON P.package_id = O.service_package_id
    WHERE O.customer_id = $1
    AND (
	    F.file_id = (
	        SELECT MIN(F2.file_id)
	        FROM services_files SF2
	        JOIN files F2 ON F2.file_id = SF2.file_id
	        WHERE SF2.service_id = S.service_id
	    ) OR F.file_id IS NULL
    )
    GROUP BY O.order_id, S.title, O.status_id, P.price, O.created_at, F.name
    ORDER BY O.created_at DESC, O.order_id
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