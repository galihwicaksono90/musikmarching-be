-- name: GetAccounts :many
select *
from account
;

-- name: GetAccountById :one
select *
from account
where id = $1
;

-- name: GetAccountByEmail :one
select *
from account
where email = $1
;

-- name: UpsertAccount :one
insert into account (email, name)
values ($1, $2)
on conflict (email) do update
set name = $2
returning id
;

