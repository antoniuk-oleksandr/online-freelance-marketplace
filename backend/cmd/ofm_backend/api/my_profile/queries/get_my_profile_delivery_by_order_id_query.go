package queries

const GetMyProfileDeliveryByOrderIdQuery = `
SELECT
    O.status_id,
    CASE
	    WHEN O.status_id = 3 OR O.status_id = 2 THEN
            json_build_object(
                'message', OD.message,
                'date', COALESCE(
                    TO_CHAR(OD.delivered_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'),
                    CASE
                        WHEN O.status_id = 2 THEN
                            TO_CHAR(
                                (O.accepted_at + make_interval(days := P.delivery_days)) AT TIME ZONE 'UTC',
                                'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'
                            )
                    END
                ),
                'files', (
                    SELECT COALESCE(
                        json_agg(F.name ORDER BY F.file_id),
                        '[]'::json
                    )
                    FROM orders_delivery_files ODF
                    JOIN files F USING (file_id)
                    WHERE ODF.delivery_id = OD.delivery_id
                )
            )
    END AS delivery,
    CASE
        WHEN O.status_id = 2 THEN
            json_build_object(
                'id', U.user_id,
                'username', U.username
	        ) 
    END AS freelancer,
    CASE
        WHEN O.status_id = 4 THEN
            json_build_object(
                'paidAmount', ROUND((P.price * (1 + $3::float))::numeric, 2)
            )
    END AS payment,
    CASE
        WHEN O.status_id = 4 THEN
            json_build_object(
                'cancellationReason', O.cancellation_reason,
                'cancelledAt', TO_CHAR(O.ended_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"')
            )
    END AS cancellation
FROM orders O
LEFT JOIN orders_deliveries OD ON
    OD.order_id = O.order_id
LEFT JOIN users U ON
    U.user_id = O.freelancer_id
LEFT JOIN packages P ON
    P.package_id = O.service_package_id
WHERE
    O.order_id = $1 AND
    (O.freelancer_id = $2 OR O.customer_id = $2)
`