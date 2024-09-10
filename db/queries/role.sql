-- name: CreateRole :one
INSERT INTO Role (
  name
) VALUES (
  $1
) RETURNING *;

-- name: GetRoleByName :one
SELECT * FROM Role
WHERE name = $1;
