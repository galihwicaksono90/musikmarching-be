-- name: GetRoleByName :one
select *
from role 
where name = $1
;

