package queries

const GetMyProfileRequirementsByOrderIDQuery = `
SELECT
    SQ.content AS question,
    OA.content AS answer
FROM orders O
LEFT JOIN service_questions SQ ON SQ.service_id = O.service_id
LEFT JOIN order_answers OA ON OA.service_question_id = SQ.service_question_id AND OA.order_id = O.order_id
WHERE O.order_id = $1 AND (O.freelancer_id = $2 OR O.customer_id = $2);
`