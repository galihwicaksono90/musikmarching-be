CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" int PRIMARY KEY,
  "title" varchar NOT NULL,
  "body" text,
  "user_id" int,
  "status" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

create table "oauth_access_tokens" (
  "id" INT PRIMARY KEY,
  "oauth_client_id" INT NULL,
  "user_id" INT NOT NULL,
  "token" VARCHAR UNIQUE NULL,
  "scope" VARCHAR NULL,
  "expired_at" TIMESTAMPTZ NULL,
  "created_by" INT NULL,
  "updated_by" INT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMPTZ NULL,
  "deleted_at" TIMESTAMPTZ NULL
);

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "oauth_access_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
