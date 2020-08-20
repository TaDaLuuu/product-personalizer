-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "users" (
  "user_id" uuid PRIMARY KEY,
  "full_name" text,
  "email" text UNIQUE,
  "password" text,
  "role" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "uuid-ossp";