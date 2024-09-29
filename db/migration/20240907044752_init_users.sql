-- +goose Up
-- +goose StatementBegin
CREATE TABLE Account (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,

  UNIQUE(email)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE Account cascade;
-- +goose StatementEnd


