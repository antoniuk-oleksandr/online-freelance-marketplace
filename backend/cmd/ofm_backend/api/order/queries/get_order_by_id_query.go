package queries

const GetOrderByIdQuery = `
WITH order_data AS (
    SELECT order_id, created_at, status_id, service_id, service_package_id
    FROM orders
    WHERE order_id = $1
),
service_questions AS (
    SELECT service_question_id, content
    FROM service_questions
    WHERE service_id = (SELECT service_id FROM order_data)
),
service_data AS (
    SELECT S.service_id, S.title,
        json_build_object('package_id', P.package_id, 'price', P.price, 'title', P.title) as package,
        (
            SELECT F.name FROM services_files SF
            LEFT JOIN files F ON F.file_id = SF.file_id
            WHERE SF.service_id = S.service_id
            ORDER BY F.file_id
            LIMIT 1
        ) as image
    FROM services S
    INNER JOIN packages P ON P.package_id = (SELECT service_package_id FROM order_data)
    WHERE S.service_id = (SELECT service_id FROM order_data)
)

SELECT
    (SELECT json_build_object(
        'order_id', order_id,
        'created_at', to_char(created_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'),
        'status_id', status_id
    ) FROM order_data) AS "order",
    (SELECT json_build_object(
        'service_id', service_id,
        'title', title,
        'image', image,
        'package', package
    ) FROM service_data) AS "service",
    (SELECT json_agg(json_build_object(
        'service_question_id', service_question_id,
        'content', content
    )
    ) AS service_questions FROM service_questions)
`