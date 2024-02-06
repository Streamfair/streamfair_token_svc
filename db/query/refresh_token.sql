-- name: CreateRefreshToken :one
INSERT INTO "token_svc"."RefreshTokens" (
    user_id,
    token,
    expires_at
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetRefreshTokenByID :one
SELECT * FROM "token_svc"."RefreshTokens" WHERE id = $1 LIMIT 1;

-- name: GetRefreshTokenByValue :one
SELECT * FROM "token_svc"."RefreshTokens" WHERE token = $1 LIMIT 1;

-- name: ListRefreshTokens :many
SELECT * FROM "token_svc"."RefreshTokens" ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateRefreshToken :one
UPDATE "token_svc"."RefreshTokens" 
SET 
    user_id = COALESCE($2, user_id),
    token = COALESCE($3, token),
    expires_at = COALESCE($4, expires_at)
WHERE id = $1 RETURNING token, expires_at;

-- name: DeleteRefreshToken :exec
DELETE FROM "token_svc"."RefreshTokens" WHERE id = $1;

-- name: VerifyRefreshToken :one
SELECT * FROM "token_svc"."RefreshTokens"
WHERE token = $1 AND revoked = false AND expires_at > NOW() LIMIT 1;

-- name: RevokeRefreshToken :exec
UPDATE "token_svc"."RefreshTokens" SET revoked = true WHERE id = $1;

-- name: GetRevokedRefreshTokens :many
SELECT id, user_id, token, revoked, expires_at, created_at
FROM "token_svc"."RefreshTokens"
WHERE revoked = true ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: BlacklistRefreshToken :exec
UPDATE "token_svc"."RefreshTokens" SET revoked = true WHERE token = $1;
