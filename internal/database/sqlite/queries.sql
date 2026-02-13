-- name: CreateUser :one
INSERT INTO users (id, credit)
VALUES (?, ?)
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ? AND is_deleted = 0
LIMIT 1;

-- name: GetAllUsers :many
SELECT *
FROM users;

-- name: DeleteUser :exec
UPDATE users
SET is_deleted = 1
WHERE id = ?;

-- name: UpdateUserByID :one
UPDATE users
SET
    is_banned   = ?,
    is_premium  = ?,
    is_verified = ?,
    total_links = ?
WHERE id = ?
RETURNING *;

-- name: GetTotalActiveUsersCount :one
SELECT COUNT(*)
FROM users
WHERE is_deleted = 0;

-- name: IncrementCreditWithDate :one
UPDATE users
SET credit = credit + ?,
    last_credit_update = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING *;

-- name: IncrementCredit :one
UPDATE users
SET credit = credit + ?
WHERE id = ?
RETURNING *;

-- name: DecrementCredit :one
UPDATE users
SET credit = credit - ?
WHERE id = ?
RETURNING *;

-- name: IncrementTotalLinks :one
UPDATE users SET total_links = total_links + 1 WHERE id = ? RETURNING *;
