-- name: ListUrls :many
SELECT * FROM urls ORDER BY filename;

-- name: CreateUrl :one
INSERT INTO urls (id, filename, url, path) VALUES ($1, $2, $3, $4) RETURNING id, filename, url;