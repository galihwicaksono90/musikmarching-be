-- name: CreateUser :one
insert into UserAccount (
  name, email, password, roleId
) values (
  $1, $2, $3, $4
) returning id, name, email, created_at
;

-- name: FindUsers :many
select u.id, u.name, u.email, u.created_at, r.name as rolename
from useraccount as u
inner join role as r on u.roleid = r.id
limit 10
;

-- name: FindUserByEmail :one
select u.id, u.name, u.email, u.created_at, u.password, r.name as rolename
from useraccount as u
inner join role as r on u.roleid = r.id
where u.email = $1
limit 1
;

