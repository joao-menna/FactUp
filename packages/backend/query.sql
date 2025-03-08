-- ########## USERS ##########

-- name: FindUserById :one
SELECT * FROM "user"
WHERE id = $1
LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM "user"
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user";

-- name: InsertUser :one
INSERT INTO "user" (email, display_name, image_path, category)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUser :exec
UPDATE "user"
SET display_name = $1,
    image_path = $2,
    category = $3
WHERE id = $4;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;



-- ########## POSTS ##########

-- name: FindPostById :one
SELECT * FROM "post"
WHERE id = $1
LIMIT 1;

-- name: FindRandomPosts :many
SELECT * FROM "post"
WHERE "type" = $1
ORDER BY RANDOM()
LIMIT $2;

-- name: FindPostsByUserId :many
SELECT * FROM "post"
WHERE user_id = $1
LIMIT $2
OFFSET $3;

-- name: InsertPost :one
INSERT INTO "post" ("type", user_id, body, source, image_path)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeletePostById :exec
DELETE FROM "post"
WHERE id = $1;



-- ########## INTERACTIONS ##########

-- name: InsertUserInteraction :one
INSERT INTO "user_interaction" (post_id, user_id, score)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteUserInteraction :exec
DELETE FROM "user_interaction"
WHERE post_id = $1 AND user_id = $2;
