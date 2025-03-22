package queries

const GetMyProfileRequestsQuery = `
WITH requests_data AS (
    SELECT
        O.order_id,
        S.title,
        O.status_id AS status,
        (
            SELECT P.price FROM packages P
            WHERE P.service_id = S.service_id
            ORDER BY P.package_id
            LIMIT 1
        ) AS price,
        U.first_name AS customer_first_name,
        U.surname AS customer_surname,
        COALESCE(F.name, '') as customer_avatar,
        TO_CHAR(S.created_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"') AS date
    FROM orders O
    LEFT JOIN services S ON S.service_id = O.service_id
    LEFT JOIN users U ON U.user_id = O.customer_id
    LEFT JOIN files F ON f.file_id = U.avatar_id
    WHERE O.freelancer_id = $1
    AND (
        CASE 
            WHEN $2 = 0 THEN O.status_id IN (2, 3, 4, 7) 
            ELSE O.status_id = $2 
        END
    )
    ORDER BY O.created_at DESC, O.order_id
    OFFSET $3
    LIMIT $4
),
total_pages_data AS (
    SELECT CEIL(CAST(count(O.order_id) AS FLOAT) / $4) as total_pages
    FROM orders O
    WHERE freelancer_id = $1
    AND (
        CASE 
            WHEN $2 = 0 THEN O.status_id IN (2, 3, 4, 7) 
            ELSE O.status_id = $2 
        END
    )
)

SELECT
COALESCE((
    SELECT json_agg(json_build_object(
        'order_id', RD.order_id,
        'title', RD.title,
        'status', RD.status,
        'price', RD.price,
        'customer_first_name', RD.customer_first_name,
        'customer_surname', RD.customer_surname,
        'customer_avatar', RD.customer_avatar,
        'date', RD.date
    )) FROM requests_data RD
), '[]'::json) AS requests_table_data,
(
    SELECT * FROM total_pages_data
)
`