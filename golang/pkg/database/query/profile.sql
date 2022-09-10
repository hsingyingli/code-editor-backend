-- name: GetUser :one
SELECT * FROM profile
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO profile (
  username, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM profile
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE profile
set username = $2,
email = $3,
password = $4,
avatar_url = $5
WHERE id = $1;
