-- name: OauthFindOneByAccessToken :one
select * from oauth_access_tokens where token = $1 limit 1;

-- name: OauthCreate :one
INSERT INTO oauth_access_tokens (
  user_id,
  token
) VALUES (
  $1, $2
) RETURNING *;
