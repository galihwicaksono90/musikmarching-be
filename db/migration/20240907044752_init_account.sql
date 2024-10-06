-- +goose Up
-- +goose StatementBegin
CREATE TYPE RoleName AS ENUM ('admin', 'contributor', 'user');

CREATE TABLE Role (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name RoleName NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,

  UNIQUE(name)
);

insert into role as r (id, name)
values 
  ('515ebadd-f8d1-472f-8b8c-1ba53d61a358', 'admin'),
  ('78653d16-6134-4f84-afb6-f44deb51f898', 'contributor'),
  ('748d4130-fc58-4485-be80-f49342252132', 'user')
;

CREATE TABLE Account (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  pictureUrl VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,
  role_id UUID NOT NULL REFERENCES Role(id),

  UNIQUE(email)
);

CREATE TABLE Profile (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  account_id UUID NOT NULL REFERENCES Account(id),

  UNIQUE(account_id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE Role cascade;
DROP TABLE Account cascade;
DROP TABLE Profile cascade;
DROP type RoleName cascade;
-- +goose StatementEnd


