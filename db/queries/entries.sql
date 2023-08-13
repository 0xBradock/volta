-- name: GetEntryByID :one
SELECT * FROM entry
WHERE id = $1 LIMIT 1;

-- name: GetEntriesByDomain :many
SELECT * FROM entry
WHERE domain LIKE CONCAT('%', @domainSubStr::text, '%');

-- name: GetEntries :many
SELECT * FROM entry
ORDER BY domain
LIMIT $1
OFFSET $2;

-- name: CreateEntry :one
INSERT INTO entry (
  id,
  domain,
  login,
  password,
  meta
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateEntry :one
UPDATE entry
SET
  login = $2,
  password = $3,
  meta = $4,
  updated = now()
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :one
DELETE FROM entry
WHERE id = $1
RETURNING *;
