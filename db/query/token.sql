-- name: CreateToken :one
INSERT INTO "token_svc"."Tokens" (
    user_id,
    token_type,
    token,
    expires_at
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetTokenByID :one
SELECT * FROM "token_svc"."Tokens" WHERE id = $1 LIMIT 1;

-- name: GetTokenByValue :one
SELECT * FROM "token_svc"."Tokens" WHERE token = $1 LIMIT 1;

-- name: ListTokens :many
SELECT * FROM "token_svc"."Tokens" ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateToken :one
UPDATE "token_svc"."Tokens" 
SET
    user_id = COALESCE(sqlc.narg(user_id), user_id),
    token_type = COALESCE(sqlc.narg(token_type), token_type),
    updated_at = now()
WHERE id = sqlc.arg(id) RETURNING *;

-- name: VerifyToken :one
SELECT * FROM "token_svc"."Tokens" WHERE token = $1
AND revoked = false AND expires_at > NOW() LIMIT 1;

-- name: RevokeTokenByID :exec
UPDATE "token_svc"."Tokens" SET revoked = true WHERE id = $1;

-- name: RevokeTokenByValue :exec
UPDATE "token_svc"."Tokens" SET revoked = true WHERE token = $1;

-- name: DeleteToken :exec
DELETE FROM "token_svc"."Tokens" WHERE id = $1;

-- name: DeleteTokenByValue :exec
DELETE FROM "token_svc"."Tokens" WHERE token = $1;

-- name: ListRevokedTokens :many
SELECT *
FROM "token_svc"."Tokens"
WHERE revoked = true ORDER BY created_at DESC LIMIT $1 OFFSET $2;

