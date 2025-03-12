package utils

const UserByIdQuery = `
SELECT
    U.user_id, U.username, U.about, U.created_at,
    U.first_name, U.level, U.surname, F.name AS avatar,
    -- Languages subquery
    (
        SELECT json_agg(json_build_object('language_id', LI.language_id, 'name', LI.name))
        FROM languages LI
            JOIN users_languages ULI ON ULI.language_id = LI.language_id
        WHERE ULI.user_id = U.user_id
    ) AS languages,
    -- User skills subquery
    (
        SELECT json_agg(json_build_object('skill_id', SI.skill_id, 'name', SI.name))
        FROM skills SI
            JOIN users_skills USI ON USI.skill_id = SI.skill_id
        WHERE USI.user_id = U.user_id
    )  AS skills,
    COALESCE(userCountRating.reviews_count, 0) as reviews_count,
    COALESCE(ROUND(userCountRating.rating, 2), 0) as rating
FROM users U
    LEFT JOIN files F ON F.file_id = U.avatar_id
    -- User count rating
    LEFT JOIN LATERAL (
        SELECT COUNT(R.review_id) as reviews_count, AVG(R.rating) as rating
            FROM orders O
                LEFT JOIN services S ON S.service_id = O.service_id
                LEFT JOIN reviews R on R.review_id = O.review_id
            WHERE S.freelancer_id = U.user_id AND R.review_id IS NOT NULL
    ) AS userCountRating ON true
WHERE U.user_id = $1;
`

const ReviewsByUserIdQuery = `
WITH
    reviews_data AS (
    SELECT
        DISTINCT ON (R.review_id)
        R.review_id AS review_id, R.content, R.rating,
        TO_CHAR(O.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.MS+00:00') AS created_at,
        TO_CHAR(O.ended_at, 'YYYY-MM-DD"T"HH24:MI:SS.MS+00:00') AS ended_at,
        U.user_id AS user_id, U.username,
        F.name AS avatar, S.service_id AS service_id,
        P.price, F1.name AS service_image,
        S.title
    FROM orders O
    JOIN reviews R ON O.review_id = R.review_id
    LEFT JOIN users U ON U.user_id = O.customer_id
    LEFT JOIN files F ON U.avatar_id = F.file_id
    LEFT JOIN services S ON O.service_id = S.service_id
    LEFT JOIN packages P ON P.package_id = O.service_package_id
    LEFT JOIN services_files SF ON S.service_id = SF.service_id
    LEFT JOIN files F1 ON F1.file_id = SF.file_id
    WHERE 
    O.freelancer_id = $1
    AND(
        ($2 = '' OR $2 IS NULL OR $3 = -1 OR $3 IS NULL)
        OR
        (
            (TO_CHAR(O.ended_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS.MS+00:00') = TO_CHAR($2::timestamp AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS.MS+00:00') AND R.review_id < $3)
            OR
            ended_at < $2::timestamp
        )
    )
    ),
    ranked_reviews_ordered AS (
        SELECT * FROM reviews_data AS RV
        ORDER BY RV.ended_at DESC, RV.review_id DESC
        LIMIT $4
    )
SELECT
    (SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)) AS user_exists,
    jsonb_agg(
        jsonb_build_object(
            'review_id', ranked_reviews_ordered.review_id,
            'content', ranked_reviews_ordered.content,
            'rating', ranked_reviews_ordered.rating,
            'created_at', ranked_reviews_ordered.created_at,
            'ended_at', ranked_reviews_ordered.ended_at,
            'user_id', ranked_reviews_ordered.user_id,
            'username', ranked_reviews_ordered.username,
            'avatar', ranked_reviews_ordered.avatar,
            'service_id', ranked_reviews_ordered.service_id,
            'price', ranked_reviews_ordered.price,
            'service_image', ranked_reviews_ordered.service_image,
            'title', ranked_reviews_ordered.title
        )
    ) AS data
FROM ranked_reviews_ordered;
`

const ServicesByUserIdQuery = `
WITH services_data AS (
    SELECT
        DISTINCT ON (S.service_id)
        S.service_id AS services_data_id,
        TO_CHAR(S.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.MS+00:00') AS created_at,
        S.description,
        S.title, S.category_id, S.freelancer_id,
        F.name AS image,
        COALESCE(subCountRating.count, 0) AS reviews_count,
        COALESCE(subCountRating.rating, 0) AS rating,
        COALESCE(subMinPrice.minPrice, 0) AS min_price
    FROM services S
    LEFT JOIN services_files SF ON S.service_id = SF.service_id
    LEFT JOIN files F ON SF.file_id = F.file_id
    LEFT JOIN (
        SELECT
            S.service_id,
            COUNT(R.review_id) AS count,
            ROUND(AVG(R.rating), 2) AS rating
        FROM services S
            LEFT JOIN orders O ON S.service_id = O.service_id
            LEFT JOIN reviews R ON O.review_id = R.review_id
        WHERE R.review_id IS NOT NULL
        GROUP BY S.service_id
    ) subCountRating ON S.service_id = subCountRating.service_id
    LEFT JOIN (
        SELECT
            SP.service_id AS id, MIN(SP.price) as minPrice
        FROM packages SP
        GROUP BY SP.service_id
    ) subMinPrice ON S.service_id = subMinPrice.id
    WHERE S.freelancer_id = $1
    AND (
	    ($2 = -1 OR $2 IS NULL OR $3 = -1 OR $3 IS NULL)
	    OR
	    (
	    	(COALESCE(subCountRating.count, 0) = $2 AND S.service_id < $3)
	    	OR
	    	COALESCE(subCountRating.count, 0) < $2
	    )
    )
),
ranked_services_ordered AS (
    SELECT * FROM services_data AS SD
    ORDER BY SD.reviews_count DESC, SD.services_data_id DESC
    LIMIT $4
)
SELECT
    (SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)) AS user_exists,
    jsonb_agg(
        jsonb_build_object(
                'service_id', ranked_services_ordered.services_data_id,
                'created_at', ranked_services_ordered.created_at,
                'description', ranked_services_ordered.description,
                'title', ranked_services_ordered.title,
                'category_id', ranked_services_ordered.category_id,
                'freelancer_id', ranked_services_ordered.freelancer_id,
                'image', ranked_services_ordered.image,
                'reviews_count', ranked_services_ordered.reviews_count,
                'rating', ranked_services_ordered.rating,
                'min_price', ranked_services_ordered.min_price
        )
    ) as data
FROM ranked_services_ordered;
`