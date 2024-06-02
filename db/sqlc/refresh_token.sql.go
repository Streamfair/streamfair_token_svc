// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: refresh_token.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRefreshToken = `-- name: CreateRefreshToken :one
INSERT INTO "token_svc"."RefreshTokens" (
    user_id,
    token,
    expires_at
) VALUES (
    $1, $2, $3
) RETURNING id, user_id, token, revoked, expires_at, created_at, updated_at
`

type CreateRefreshTokenParams struct {
	UserID    int64     `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (q *Queries) CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (TokenSvcRefreshToken, error) {
	row := q.db.QueryRow(ctx, createRefreshToken, arg.UserID, arg.Token, arg.ExpiresAt)
	var i TokenSvcRefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.Revoked,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRefreshTokenById = `-- name: DeleteRefreshTokenById :exec
DELETE FROM "token_svc"."RefreshTokens" WHERE id = $1
`

func (q *Queries) DeleteRefreshTokenById(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteRefreshTokenById, id)
	return err
}

const deleteRefreshTokenByValue = `-- name: DeleteRefreshTokenByValue :exec
DELETE FROM "token_svc"."RefreshTokens" WHERE token = $1
`

func (q *Queries) DeleteRefreshTokenByValue(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, deleteRefreshTokenByValue, token)
	return err
}

const getRefreshTokenById = `-- name: GetRefreshTokenById :one
SELECT id, user_id, token, revoked, expires_at, created_at, updated_at FROM "token_svc"."RefreshTokens" WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRefreshTokenById(ctx context.Context, id int64) (TokenSvcRefreshToken, error) {
	row := q.db.QueryRow(ctx, getRefreshTokenById, id)
	var i TokenSvcRefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.Revoked,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRefreshTokenByValue = `-- name: GetRefreshTokenByValue :one
SELECT id, user_id, token, revoked, expires_at, created_at, updated_at FROM "token_svc"."RefreshTokens" WHERE token = $1 LIMIT 1
`

func (q *Queries) GetRefreshTokenByValue(ctx context.Context, token string) (TokenSvcRefreshToken, error) {
	row := q.db.QueryRow(ctx, getRefreshTokenByValue, token)
	var i TokenSvcRefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.Revoked,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listRefreshTokens = `-- name: ListRefreshTokens :many
SELECT id, user_id, token, revoked, expires_at, created_at, updated_at FROM "token_svc"."RefreshTokens" ORDER BY id LIMIT $1 OFFSET $2
`

type ListRefreshTokensParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRefreshTokens(ctx context.Context, arg ListRefreshTokensParams) ([]TokenSvcRefreshToken, error) {
	rows, err := q.db.Query(ctx, listRefreshTokens, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TokenSvcRefreshToken{}
	for rows.Next() {
		var i TokenSvcRefreshToken
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Token,
			&i.Revoked,
			&i.ExpiresAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRevokedRefreshTokens = `-- name: ListRevokedRefreshTokens :many
SELECT id, user_id, token, revoked, expires_at, created_at, updated_at
FROM "token_svc"."RefreshTokens"
WHERE revoked = true ORDER BY created_at DESC LIMIT $1 OFFSET $2
`

type ListRevokedRefreshTokensParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRevokedRefreshTokens(ctx context.Context, arg ListRevokedRefreshTokensParams) ([]TokenSvcRefreshToken, error) {
	rows, err := q.db.Query(ctx, listRevokedRefreshTokens, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TokenSvcRefreshToken{}
	for rows.Next() {
		var i TokenSvcRefreshToken
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Token,
			&i.Revoked,
			&i.ExpiresAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const revokeRefreshTokenById = `-- name: RevokeRefreshTokenById :exec
UPDATE "token_svc"."RefreshTokens" SET revoked = true WHERE id = $1
`

func (q *Queries) RevokeRefreshTokenById(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, revokeRefreshTokenById, id)
	return err
}

const revokeRefreshTokenByValue = `-- name: RevokeRefreshTokenByValue :exec
UPDATE "token_svc"."RefreshTokens" SET revoked = true WHERE token = $1
`

func (q *Queries) RevokeRefreshTokenByValue(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, revokeRefreshTokenByValue, token)
	return err
}

const updateRefreshToken = `-- name: UpdateRefreshToken :one
UPDATE "token_svc"."RefreshTokens" 
SET 
    user_id = COALESCE($1, user_id),
    updated_at = now()
WHERE id = $2 RETURNING id, user_id, token, revoked, expires_at, created_at, updated_at
`

type UpdateRefreshTokenParams struct {
	UserID pgtype.Int8 `json:"user_id"`
	ID     int64       `json:"id"`
}

func (q *Queries) UpdateRefreshToken(ctx context.Context, arg UpdateRefreshTokenParams) (TokenSvcRefreshToken, error) {
	row := q.db.QueryRow(ctx, updateRefreshToken, arg.UserID, arg.ID)
	var i TokenSvcRefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.Revoked,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const verifyRefreshToken = `-- name: VerifyRefreshToken :one
SELECT id, user_id, token, revoked, expires_at, created_at, updated_at FROM "token_svc"."RefreshTokens"
WHERE token = $1 AND revoked = false AND expires_at > NOW() LIMIT 1
`

func (q *Queries) VerifyRefreshToken(ctx context.Context, token string) (TokenSvcRefreshToken, error) {
	row := q.db.QueryRow(ctx, verifyRefreshToken, token)
	var i TokenSvcRefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.Revoked,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
