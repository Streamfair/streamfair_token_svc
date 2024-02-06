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
    user_id = COALESCE($2, user_id),
    token_type = COALESCE($3, token_type),
    token = COALESCE($4, token),
    expires_at = COALESCE($5, expires_at)
WHERE id = $1 RETURNING token, expires_at;

-- name: VerifyToken :one
SELECT * FROM "token_svc"."Tokens" WHERE token = $1
AND revoked = false AND expires_at > NOW() LIMIT 1;

-- name: RevokeToken :exec
UPDATE "token_svc"."Tokens" SET revoked = true WHERE id = $1;

-- name: DeleteToken :exec
DELETE FROM "token_svc"."Tokens" WHERE id = $1;

-- name: GetRevokedTokens :many
SELECT id, user_id, token_type, token, revoked, expires_at, created_at
FROM "token_svc"."Tokens"
WHERE revoked = true ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: BlacklistToken :exec
UPDATE "token_svc"."Tokens" SET revoked = true WHERE token = $1;
