package queries

const AddOrderAnswersQuery = `
INSERT INTO order_answers 
(service_question_id, content, order_id)
VALUES (:service_question_id, :content, :order_id);
`