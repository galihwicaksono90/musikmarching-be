-- name: GetAccounts :many
select 
  a.id,
  a.name,
  a.email,
  a.pictureurl,
  r.name as role_name,
  p.account_id
from account as a
inner join role as r on a.role_id = r.id
inner join profile as p on a.id = p.account_id
;

-- name: GetAccountById :one
select 
  a.id,
  a.name,
  a.email,
  a.pictureurl,
  r.name as role_name,
  p.account_id
from account as a
inner join role as r on a.role_id = r.id
inner join profile as p on a.id = p.account_id
where a.id = $1
limit 1
;

-- name: GetAccountByEmail :one
select 
  a.id,
  a.name,
  a.email,
  a.pictureurl,
  r.name as role_name,
  p.account_id
from account as a
inner join role as r on a.role_id = r.id
inner join profile as p on a.id = p.account_id
where a.email = $1
limit 1
;

-- name: CreateAccountOld :one
insert into account as a (email, name, pictureurl, role_id)
values (@email, @name, @pictureurl, @roleId)
returning id
;

-- name: UpdateAccount :one
update account as a
set 
  name = coalesce(sqlc.narg('name'), a.name),
  pictureurl = coalesce(sqlc.narg('pictureurl'), a.pictureurl)
where id = @id
returning id
;

-- name: CreateAccount :one
WITH account_insert AS (
  INSERT INTO Account (email, name, pictureUrl, role_id)
  VALUES (@email, @name, @pictureurl, @roleId)
  RETURNING id
)
INSERT INTO Profile as p (account_id)
SELECT id FROM account_insert
returning account_id
;

