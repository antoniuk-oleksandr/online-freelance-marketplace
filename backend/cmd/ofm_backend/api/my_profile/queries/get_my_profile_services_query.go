package queries

const GetMyProfileServicesQuery = `
WITH services_data AS (
    SELECT
        S.service_id, S.title,
        COALESCE ( (
            SELECT P.price FROM packages P
            WHERE P.service_id = S.service_id
            ORDER BY P.price
            LIMIT 1
        ), 0)as price,
        C.name as category,
        ROUND(COALESCE(AVG(R.rating), 0), 2) AS rating,
        COUNT(O.order_id) AS orders_count,
        TO_CHAR(S.created_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"') AS date,
        COALESCE(F.name, '') AS image
    FROM services S
    LEFT JOIN categories C ON C.category_id = S.category_id
    LEFT JOIN orders O ON o.service_id = S.service_id
    LEFT JOIN reviews R ON R.review_id = O.review_id
    LEFT JOIN services_files SF ON SF.service_id = S.service_id
    LEFT JOIN files F ON F.file_id = SF.file_id
    WHERE S.freelancer_id = $1
    AND (
        F.file_id = (
            SELECT MIN(F2.file_id)
            FROM services_files SF2
            JOIN files F2 ON F2.file_id = SF2.file_id
            WHERE SF2.service_id = S.service_id
        ) OR F.file_id IS NULL
    )
    GROUP BY S.created_at, S.service_id, C.name, F.name
    ORDER BY S.created_at DESC, S.service_id
    OFFSET $2
    LIMIT $3
),
total_pages_data AS (
    SELECT CEIL(CAST(count(S.service_id) AS FLOAT) / $3) as total_pages
    FROM services S
    WHERE freelancer_id = $1
)

SELECT
COALESCE ((
    SELECT json_agg(json_build_object(
            'service_id', SF.service_id,
            'title', SF.title,
            'price', SF.price,
            'category', SF.category,
            'rating', SF.rating,
            'orders_count', SF.orders_count,
            'date', SF.date,
            'image', SF.image
        ))
    FROM services_data SF
), '[]'::json) as services_data,
(
    SELECT * FROM total_pages_data
)
`