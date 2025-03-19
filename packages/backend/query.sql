-- ########## USERS ##########

-- name: FindUserById :one
SELECT *
FROM "user"
WHERE id = $1
LIMIT 1;

-- name: FindUserByProviderUserId :one
SELECT *
FROM "user"
WHERE provider_user_id = $1 AND
      "provider" = $2
LIMIT 1;

-- name: FindUserByEmail :one
SELECT *
FROM "user"
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM "user";

-- name: InsertUser :one
INSERT INTO "user" (provider_user_id, "provider", email, display_name, image_path, category)
VALUES ($1, $2, $3, $4, $5, $6)
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



-- ########## MOD ##########

-- name: BanUser :exec
UPDATE "user"
SET banned = TRUE
WHERE id = $1;

-- name: FindBotById :one
SELECT *
FROM "user_bot"
WHERE id = $1
LIMIT 1;

-- name: InsertBot :one
INSERT INTO "user_bot" ("user_id", "name", "secret")
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateBotSecret :exec
UPDATE "user_bot"
SET "secret" = $1
WHERE id = $2;

-- name: InsertBotUser :one
INSERT INTO "user" (display_name, category)
VALUES ($1, $2)
RETURNING *;



-- ########## POSTS ##########

-- name: GetPostedCountByDay :one
SELECT COUNT(*) AS postCount
FROM "post"
WHERE DATE("created_at") = CURRENT_DATE AND user_id = $1;

-- name: FindPostById :one
SELECT *
FROM "post"
WHERE id = $1
LIMIT 1;

-- name: FindRandomPosts :many
SELECT *
FROM "post"
WHERE "type" = $1
ORDER BY RANDOM()
LIMIT $2;

-- name: FindPostsByUserId :many
SELECT *
FROM "post"
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

-- name: GetInteractionScoreByPostId :one
SELECT SUM("score") AS totalScore
FROM "user_interaction"
WHERE post_id = $1;

-- name: InsertUserInteraction :one
INSERT INTO "user_interaction" (post_id, user_id, score)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteUserInteraction :exec
DELETE FROM "user_interaction"
WHERE post_id = $1 AND user_id = $2;
