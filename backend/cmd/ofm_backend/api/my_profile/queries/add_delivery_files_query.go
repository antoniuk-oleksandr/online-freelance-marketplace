package queries

const AddDeliveryFilesQuery = `
INSERT INTO orders_delivery_files
(delivery_id, file_id)
VALUES (:delivery_id, :file_id)
`