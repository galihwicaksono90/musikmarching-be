-- +goose Up
-- +goose StatementBegin
CREATE TYPE RoleType AS ENUM ('admin', 'user', 'contributor');

CREATE TABLE Role(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name RoleType NOT NULL,
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,
  UNIQUE(name)
);

CREATE TABLE UserAccount (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  roleId UUID NOT NULL,
  password VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,

  UNIQUE(email)
);

CREATE TABLE Score (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,

  uploadedById UUID NOT NULL,

  scoreType VARCHAR(255) NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,

  verifiedbyId UUID
);

ALTER TABLE UserAccount ADD FOREIGN KEY (roleId) REFERENCES Role(id);

ALTER TABLE Score ADD FOREIGN KEY (uploadedById) REFERENCES UserAccount(id);

ALTER TABLE Score ADD FOREIGN KEY (verifiedById) REFERENCES UserAccount(id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE UserAccount cascade;
DROP TABLE Role cascade;
DROP TABLE Score cascade;
DROP Type RoleType;
-- +goose StatementEnd


