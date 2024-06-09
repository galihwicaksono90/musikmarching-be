-- name: GetUsers :many
select * from users limit 10;

-- name: GetUserById :one
select * from users where id = $1 limit 1;

-- name: GetUserByEmail :one
select * from users where email = $1 limit 1;

-- name: CreateUser :exec
INSERT INTO users (
  username,
  email,
  password
) VALUES (
  $1, $2 , $3
);
