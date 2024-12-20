package utils

const initialSearchServicesQuery = `
SELECT
	DISTINCT (S.id),
    S.created_at, S.title,
    S.category_id, S.freelancer_id,
    (SELECT name
    FROM services_files SF
             INNER JOIN files F ON F.id = SF.file_id
    WHERE SF.service_id = S.id
    ORDER BY SF.file_id
    LIMIT 1) as image,
    COALESCE(subCountRating.count, 0) AS reviews_count,
    COALESCE(ROUND(subCountRating.rating, 2), 0) AS rating,
    COALESCE(subMinPrice.minPrice, 0) AS min_price,
    last_month_orders.last_month_completed_orders_count AS last_month_completed_orders_count,
    freelancer.level as level
FROM services S
LEFT JOIN (
    SELECT
        S.id,
        COUNT(R.id) AS count,
        AVG(R.rating) AS rating
    FROM services S
        LEFT JOIN orders O ON S.id = O.service_id
        LEFT JOIN reviews R ON O.review_id = R.id
    WHERE R.id IS NOT NULL
    GROUP BY S.id
) subCountRating ON S.id = subCountRating.id
LEFT JOIN (
    SELECT
        SP.service_id AS id,
        MIN(P.price) as minPrice,
        MIN(P.delivery_days) as delivery_days
    FROM services_packages SP
        LEFT JOIN packages P ON P.id = SP.package_id
    GROUP BY SP.service_id
    ) subMinPrice ON S.id = subMinPrice.id
    
LEFT JOIN (
    SELECT
        SER.id AS service_id,
        COUNT(O.id) AS last_month_completed_orders_count
    FROM services SER
        LEFT JOIN orders O
            ON SER.id = O.service_id
            AND O.status_id IN (3)
            AND O.ended_at >= DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
    GROUP BY SER.id
) last_month_orders ON S.id = last_month_orders.service_id    
    
INNER JOIN LATERAL (
    SELECT * FROM users U
    WHERE U.id = S.freelancer_id
) as freelancer ON S.freelancer_id = S.freelancer_id
`

// WHERE
//     title LOWER(LIKE) LOWER('%%') -- Query param
//     AND subMinPrice.minPrice >= 0 -- Price From
//     AND subMinPrice.minPrice <= 1000 -- Price To
//     AND COALESCE(ROUND(subCountRating.rating, 2), 0) >= 0 -- Rating from
//     AND COALESCE(ROUND(subCountRating.rating, 2), 0) <= 5 -- Rating to
//     AND subMinPrice.delivery_days >= 0 -- Delivery time From
//     AND subMinPrice.delivery_days <= 30 -- Delivery time To
//     AND freelancer.level >= 1 -- Level From
//     AND freelancer.level <= 5 -- Level To

//     AND S.category_id IN (1, 2, 3, 4, 5) -- Category

//     AND EXISTS (
//         SELECT 1
//         FROM users_languages UL
//         WHERE UL.user_id = freelancer.id
//         AND UL.language_id IN (1, 2, 3, 4, 5, 6) -- Language
//     )

//     AND EXISTS(
//         SELECT 1
//         FROM users_skills US
//         WHERE US.user_id = freelancer.id
//         AND US.skill_id IN (4) -- Skill
//     )

// AND
//     (S.title = 'Vagram' AND S.id > 0)
// OR
//     (S.title < 'Vagram')
// ORDER BY title DESC, id DESC -- Order | SORT param
// LIMIT 51 -- Limit param