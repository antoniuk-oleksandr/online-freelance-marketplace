package utils

const GetHomeDataQuery = `
WITH best_freelancers AS (
        WITH data as (
        SELECT
            U.user_id,
            U.first_name,
            U.surname,
            F.name as avatar,
            ROUND(COALESCE(AVG(R.rating), 0), 2) as rating,
            (
                SELECT COUNT(O.order_id) FROM orders
                O WHERE O.freelancer_id = U.user_id AND O.status_id = 3
            ) as completed_projects
        FROM users U
        LEFT JOIN files F ON f.file_id = U.avatar_id
        LEFT JOIN orders O ON O.freelancer_id = U.user_id
        LEFT JOIN reviews R ON R.review_id = O.review_id
        GROUP BY U.user_id, U.first_name, U.surname, F.name
        ORDER BY completed_projects DESC, U.user_id DESC
    )
    SELECT
        json_agg(
            json_build_object(
                'id', data.user_id,
                'first_name', data.first_name,
                'surname', data.surname,
                'avatar', data.avatar,
                'rating', data.rating,
                'completed_projects', data.completed_projects
            )
        ) AS freelances_available
    FROM data
    WHERE  data.rating > 1
    LIMIT 4
),
key_metrics AS (
    SELECT json_build_object(
        'freelances_available', COUNT(S.service_id),
        'projects_completed', (
            SELECT COUNT(O.order_id)
            FROM orders o
            WHERE o.status_id = 3
        ),
        'avg_rating', (
            SELECT ROUND(COALESCE(AVG(r.rating), 0), 2)
            FROM orders o
            LEFT JOIN reviews r ON O.review_id = o.review_id
        )
    ) as key_metrics
    FROM services S
),
best_reviews AS (
    WITH limited_reviews AS (
        SELECT U.first_name, U.surname, R.content, R.rating
        FROM orders O
        INNER JOIN reviews R ON R.review_id = O.review_id
        INNER JOIN users U ON U.user_id = O.customer_id
        WHERE R.rating >= 4
        ORDER BY LENGTH(R.content) DESC, R.review_id DESC
        LIMIT 4
    )
    SELECT json_agg(
        json_build_object(
            'first_name', first_name,
            'surname', surname,
            'content', content,
            'rating', rating
        )
    ) AS best_reviews
    FROM limited_reviews
),
best_services AS (
    WITH ranked_services AS (
        SELECT
            S.service_id,
            S.title,
            S.description,
            (
                SELECT F.name FROM services_files SF
                LEFT JOIN files F ON F.file_id = SF.file_id
                WHERE SF.service_id = S.service_id
                LIMIT 1

            ) as image,
            RANK() OVER (
                PARTITION BY S.category_id
                ORDER BY COUNT(O.order_id) DESC, S.service_id DESC
            ) AS rank
        FROM orders O
        LEFT JOIN services S ON S.service_id = O.service_id
        WHERE S.category_id IN (1, 6, 3, 2)
        AND O.status_id = 3
        GROUP BY S.service_id, S.title, S.description, S.category_id
    )
    SELECT
        json_agg(
                json_build_object(
                    'id', ranked_services.service_id,
                    'title', ranked_services.title,
                    'description', ranked_services.description,
                    'image', ranked_services.image
                )
        ) as best_services
FROM ranked_services
WHERE rank = 1
)
SELECT json_build_object(
    'best_freelancers', (SELECT freelances_available FROM best_freelancers),
    'key_metrics', (SELECT key_metrics FROM key_metrics),
    'best_freelances', (SELECT best_services FROM best_services),
    'best_reviews', (SELECT best_reviews FROM best_reviews)
) AS data;
`