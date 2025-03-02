-- name: GetUserById :one
SELECT * FROM "user"
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "user"
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user";

-- name: InsertUser :one
INSERT INTO "user" (email, display_name)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUser :exec
UPDATE "user"
SET display_name = $1
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;
