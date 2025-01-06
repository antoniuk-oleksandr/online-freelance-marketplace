package utils

const GetHomeDataQuery = `
WITH best_freelancers AS (
        WITH data as (
        SELECT
            U.id,
            U.first_name,
            U.surname,
            F.name as avatar,
            ROUND(COALESCE(AVG(R.rating), 0), 2) as rating,
            (
                SELECT COUNT(O.id) FROM orders
                O WHERE O.freelancer_id = U.id AND O.status_id = 3
            ) as completed_projects
        FROM users U
        LEFT JOIN files F ON f.id = U.avatar_id
        LEFT JOIN orders O ON O.freelancer_id = U.id
        LEFT JOIN reviews R ON R.id = O.review_id
        GROUP BY U.id, U.first_name, U.surname, F.name
        ORDER BY completed_projects DESC, U.id DESC
    )
    SELECT
        json_agg(
            json_build_object(
                'id', data.id,
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
        'freelances_available', COUNT(S.id),
        'projects_completed', (
            SELECT COUNT(o.id)
            FROM orders o
            WHERE o.status_id = 3
        ),
        'avg_rating', (
            SELECT ROUND(COALESCE(AVG(r.rating), 0), 2)
            FROM orders o
            LEFT JOIN reviews r ON r.id = o.review_id
        )
    ) as key_metrics
    FROM services S
),
best_reviews AS (
    WITH limited_reviews AS (
        SELECT U.first_name, U.surname, R.content, R.rating
        FROM orders O
        INNER JOIN reviews R ON R.id = O.review_id
        INNER JOIN users U ON U.id = O.customer_id
        WHERE R.rating >= 4
        ORDER BY LENGTH(R.content) DESC, R.id DESC
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
            S.id,
            S.title,
            S.description,
            (
                SELECT F.name FROM services_files SF
                LEFT JOIN files F ON F.id = SF.file_id
                WHERE SF.service_id = S.id
                LIMIT 1

            ) as image,
            RANK() OVER (
                PARTITION BY S.category_id
                ORDER BY COUNT(O.id) DESC, S.id DESC
            ) AS rank
        FROM orders O
        LEFT JOIN services S ON S.id = O.service_id
        WHERE S.category_id IN (1, 6, 3, 2)
        AND O.status_id = 3
        GROUP BY S.id, S.title, S.description, S.category_id
    )
    SELECT
        json_agg(
                json_build_object(
                    'id', ranked_services.id,
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