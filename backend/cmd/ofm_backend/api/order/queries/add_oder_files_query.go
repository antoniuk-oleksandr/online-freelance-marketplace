package queries

const AddOrderFilesQuery = `
INSERT INTO orders_files
(order_id, file_id)
VALUES (:order_id, :file_id);
`
