package queries

var AddFilesQuery = `
INSERT INTO files (name)
VALUES (:name)
RETURNING file_id
`