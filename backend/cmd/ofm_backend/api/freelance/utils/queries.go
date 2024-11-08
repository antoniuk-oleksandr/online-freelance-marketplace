package utils

var FreelanceQuery = `
SELECT
    S.id, S.created_at, S.description,S.title,
    countRating.count as reviews_count,
    countRating.avg as rating,
    (
        SELECT
            json_agg(FI.name)
        FROM services_files SFI
            LEFT JOIN files FI ON SFI.file_id = FI.id
        WHERE SFI.service_id = S.id
    ) as images,
    (
        SELECT
            json_build_object(
               'id', CI.id,
               'name', CI.name
            )
         FROM categories CI
         WHERE CI.id = S.category_id
    ) as category,
    (
        SELECT json_agg
            (json_build_object(
                'id', PI.id,
                'delivery_days', PI.delivery_days,
                'description', PI.description,
                'price', PI.price,
                'title', PI.title
            ))
        FROM services SI
            LEFT JOIN services_packages SPI ON SPI.service_id = SI.id
            LEFT JOIN packages PI ON PI.id = SPI.package_id
        WHERE SI.ID = S.id
    ) as packages,
    (
        SELECT
            json_build_object(
                'id', UI.id,
                'first_name', UI.first_name,
                'surname', UI.surname,
                'avatar', FI.name,
                'rating', countRating.avg,
                'level', UI.level,
                'reviews_count', countRating.count
            )
        FROM services SI
            LEFT JOIN users UI ON UI.id = SI.freelancer_id
            LEFT JOIN LATERAL
            (
	            SELECT COUNT(RI.id), AVG(RI.rating)
                FROM orders OI
                    LEFT JOIN reviews RI ON RI.id = OI.review_id
                WHERE OI.freelancer_id = SI.freelancer_id AND ended_at IS NOT NULL
            ) as countRating ON True
            LEFT JOIN files FI ON FI.id = UI.avatar_id
        WHERE SI.id = S.id
    ) as freelancer
FROM services S
LEFT JOIN LATERAL
    (
        SELECT COUNT(RI.id), AVG(RI.rating)
        FROM orders OI
            LEFT JOIN reviews RI ON RI.id = OI.review_id
        WHERE OI.service_id = S.id AND ended_at IS NOT NULL
    ) as countRating ON True
WHERE S.id = $1`

var FreealnceReviewsQuery = `
SELECT
    R.id as id,
    R.content,
    COALESCE(R.rating, 0) as rating,
    OI.created_at,
    OI.ended_at,
    json_build_object(
        'id', UI.id,
        'username', UI.username,
        'avatar', FI.name
    ) as customer,
    json_build_object(
        'price', PI.price
    ) as service
FROM orders OI
    LEFT JOIN reviews R ON R.id = OI.review_id
    LEFT JOIN users UI ON UI.id = OI.customer_id
    LEFT JOIN files FI ON UI.avatar_id = FI.id
    LEFT JOIN packages PI ON PI.id = OI.service_package_id
WHERE OI.service_id = $1 AND OI.ended_at IS NOT NULL
ORDER BY OI.ended_at DESC
`