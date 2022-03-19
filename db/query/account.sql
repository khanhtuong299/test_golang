-- name: CreateAccount :one
INSERT INTO accounts (
  account, public_key, private_key
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountByName :one
SELECT * FROM accounts
WHERE account = $1 LIMIT 1;

-- name: ListAccount :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2
;

-- name: UpdateAccount :one
UPDATE accounts SET Amount = Amount + $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;