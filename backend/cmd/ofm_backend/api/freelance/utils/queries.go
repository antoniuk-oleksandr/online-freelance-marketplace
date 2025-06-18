package utils

var FreelanceQuery = `
SELECT
    S.service_id, S.created_at, S.description,S.title,
    countRating.count as reviews_count,
    ROUND((countRating.avg)::numeric, 2) as rating,
    (
      SELECT
        COALESCE(json_agg(FI.name), '[]'::json)
      FROM services_files SFI
      LEFT JOIN files FI ON SFI.file_id = FI.file_id
      WHERE SFI.service_id = S.service_id
    ) as images,
    (
        SELECT
            json_build_object(
               'category_id', CI.category_id,
               'name', CI.name
            )
         FROM categories CI
         WHERE CI.category_id = S.category_id
    ) as category,
    (
        SELECT jsonb_agg
            (jsonb_build_object(
                'package_id', PI.package_id,
                'delivery_days', PI.delivery_days,
                'description', PI.description,
                'price', PI.price,
                'title', PI.title
            )
            ORDER BY PI.package_id
           )
        FROM packages PI
        WHERE PI.service_id = S.service_id
    ) as packages,
    (
        SELECT
            json_build_object(
                'user_id', UI.user_id,
                'username', UI.username,
                'first_name', UI.first_name,
                'surname', UI.surname,
                'avatar', FI.name,
                'rating', ROUND(countRating.avg, 2),
                'level', UI.level,
                'reviews_count', countRating.count
            )
        FROM services SI
            LEFT JOIN users UI ON UI.user_id = SI.freelancer_id
            LEFT JOIN LATERAL
            (
                SELECT COALESCE(COUNT(RI.review_id), 0) as count, COALESCE(AVG(RI.rating), 0) as avg
                FROM orders OI
                    LEFT JOIN reviews RI ON RI.review_id = OI.review_id
                WHERE OI.freelancer_id = SI.freelancer_id AND ended_at IS NOT NULL
            ) as countRating ON True
            LEFT JOIN files FI ON FI.file_id = UI.avatar_id
        WHERE SI.service_id = S.service_id
    ) as freelancer
FROM services S
LEFT JOIN LATERAL
    (
        SELECT COALESCE(COUNT(RI.review_id), 0) as count, COALESCE(AVG(RI.rating), 0) as avg
        FROM orders OI
            LEFT JOIN reviews RI ON RI.review_id = OI.review_id
        WHERE OI.service_id = S.service_id AND ended_at IS NOT NULL
    ) as countRating ON True
WHERE S.service_id = $1;
`

var FreealnceReviewsQuery = `
SELECT
    R.review_id as review_id,
    R.content,
    COALESCE(R.rating, 0) as rating,
    OI.created_at,
    OI.ended_at,
    json_build_object(
        'user_id', UI.user_id,
        'username', UI.username,
        'avatar', FI.name
    ) as customer,
    json_build_object(
        'price', PI.price
    ) as service
FROM orders OI
    LEFT JOIN reviews R ON R.review_id = OI.review_id
    LEFT JOIN users UI ON UI.user_id = OI.customer_id
    LEFT JOIN files FI ON UI.avatar_id = FI.file_id
    LEFT JOIN packages PI ON PI.package_id = OI.service_package_id
WHERE
    OI.service_id = $1
    AND OI.ended_at IS NOT NULL
    AND R.review_id IS NOT NULL
    AND (
           ($2 = '' OR $2 IS NULL OR $3 = -1 OR $3 IS NULL)
           OR (
	           (OI.ended_at = $2::timestamp AND OI.order_id < $3)
	           OR OI.ended_at < $2::timestamp
           )
    )
ORDER BY OI.ended_at DESC, OI.order_id DESC
LIMIT $4
`
var RestrictedFreelanceQuery = `
SELECT
    S.service_id, S.title,
    countRating.count AS reviews_count,
    ROUND((countRating.avg)::numeric, 2) AS rating,
    (
        SELECT FI.name
        FROM services_files SFI
        LEFT JOIN files FI ON SFI.file_id = FI.file_id
        WHERE SFI.service_id = S.service_id
        LIMIT 1
    ) AS image,
    (
        SELECT json_agg(
            json_build_object(
            	'package_id', PI.package_id,
                'delivery_days', PI.delivery_days,
                'description', PI.description,
                'price', PI.price,
                'title', PI.title
            )
            ORDER BY PI.package_id
        )
        FROM packages PI
        WHERE PI.service_id = S.service_id
    ) AS packages
FROM services S
LEFT JOIN LATERAL (
    SELECT COALESCE(COUNT(RI.review_id), 0) AS count, COALESCE(AVG(RI.rating), 0) AS avg
    FROM orders OI
    LEFT JOIN reviews RI ON RI.review_id = OI.review_id
    WHERE OI.service_id = S.service_id AND OI.ended_at IS NOT NULL
) AS countRating ON TRUE
WHERE S.service_id = $1;
`