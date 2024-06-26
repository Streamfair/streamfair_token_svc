-- name: CreateRefreshToken :one
INSERT INTO "token_svc"."RefreshTokens" (
    user_id,
    token,
    expires_at,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) RETURNING *;

-- name: GetRefreshTokenById :one
SELECT * FROM "token_svc"."RefreshTokens" WHERE id = $1 LIMIT 1;

-- name: GetRefreshTokenByValue :one
SELECT * FROM "token_svc"."RefreshTokens" WHERE token = $1 LIMIT 1;

-- name: ListRefreshTokens :many
SELECT * FROM "token_svc"."RefreshTokens" ORDER BY id LIMIT $1 OFFSET $2;

-- name: VerifyRefreshToken :one
SELECT * FROM "token_svc"."RefreshTokens"
WHERE token = $1 AND revoked = false AND expires_at > NOW() LIMIT 1;

-- name: RevokeRefreshTokenById :exec
UPDATE "token_svc"."RefreshTokens" SET revoked = true WHERE id = $1;

-- name: RevokeRefreshTokenByValue :exec
UPDATE "token_svc"."RefreshTokens" SET revoked = true WHERE token = $1;

-- name: DeleteRefreshTokenById :exec
DELETE FROM "token_svc"."RefreshTokens" WHERE id = $1;

-- name: DeleteRefreshTokenByValue :exec
DELETE FROM "token_svc"."RefreshTokens" WHERE token = $1;

-- name: ListRevokedRefreshTokens :many
SELECT *
FROM "token_svc"."RefreshTokens"
WHERE revoked = true ORDER BY created_at DESC LIMIT $1 OFFSET $2;
