package utils

const initialSearchServicesQuery = `
SELECT DISTINCT
    S.service_id,
    S.created_at,
    S.title,
    S.category_id,
    S.freelancer_id,
    (
        SELECT name
        FROM services_files SF
        INNER JOIN files F ON F.file_id = SF.file_id
        WHERE SF.service_id = S.service_id
        ORDER BY SF.file_id
        LIMIT 1
    ) AS image,
    COALESCE(subCountRating.count, 0) AS reviews_count,
    COALESCE(ROUND(subCountRating.rating, 2), 0) AS rating,
    COALESCE(subMinPrice.minPrice, 0) AS min_price,
    COALESCE(last_month_orders.last_month_completed_orders_count, 0) AS last_month_completed_orders_count,
    freelancer.level AS level
FROM services S

-- Count Ratings
LEFT JOIN (
    SELECT
        O.service_id,
        COUNT(R.review_id) AS count,
        AVG(R.rating) AS rating
    FROM orders O
    LEFT JOIN reviews R ON O.review_id = R.review_id
    WHERE R.review_id IS NOT NULL
    GROUP BY O.service_id
) subCountRating ON S.service_id = subCountRating.service_id

-- Min Price
LEFT JOIN (
    SELECT
        SP.service_id,
        MIN(SP.price) AS minPrice,
        MIN(SP.delivery_days) AS delivery_days
    FROM packages SP
    GROUP BY SP.service_id
) subMinPrice ON S.service_id = subMinPrice.service_id

-- Last Month Completed Orders
LEFT JOIN (
    SELECT
        O.service_id,
        COUNT(O.order_id) AS last_month_completed_orders_count
    FROM orders O
    WHERE O.status_id IN (3)
      AND O.ended_at >= DATE_TRUNC('month', CURRENT_DATE - INTERVAL '1 month')
    GROUP BY O.service_id
) last_month_orders ON S.service_id = last_month_orders.service_id

-- Freelancer Info
INNER JOIN LATERAL (
    SELECT U.user_id, U.level
    FROM users U
    WHERE U.user_id = S.freelancer_id
) AS freelancer ON TRUE

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