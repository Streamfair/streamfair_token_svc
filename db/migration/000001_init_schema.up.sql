CREATE SCHEMA "token_svc";

CREATE TABLE "token_svc"."Tokens" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGINT NOT NULL,
  "token" TEXT NOT NULL,
  "revoked" BOOLEAN NOT NULL DEFAULT false,
  "expires_at" TIMESTAMPTZ NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "token_svc"."RefreshTokens" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGINT NOT NULL,
  "token" TEXT NOT NULL,
  "revoked" BOOLEAN NOT NULL DEFAULT false,
  "expires_at" TIMESTAMPTZ NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX "idx_tokens_id" ON "token_svc"."Tokens" ("id");

CREATE INDEX "idx_tokens_user_id" ON "token_svc"."Tokens" ("user_id");

CREATE INDEX "idx_refresh_tokens_id" ON "token_svc"."RefreshTokens" ("id");

CREATE INDEX "idx_refresh_tokens_user_id" ON "token_svc"."RefreshTokens" ("user_id");

CREATE INDEX "idx_tokens_created_at" ON "token_svc"."Tokens" ("created_at");

CREATE INDEX "idx_refresh_tokens_created_at" ON "token_svc"."RefreshTokens" ("created_at");