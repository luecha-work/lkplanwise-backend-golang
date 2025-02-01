-- name: GetAllTransactions :many
SELECT * FROM "TransactionHistory";

-- name: GetTransactionById :one
SELECT * FROM "TransactionHistory" WHERE "Id" = $1;

-- name: CreateTransaction :one
INSERT INTO "TransactionHistory" ("Id", "AccountId", "TransactionType", "Amount", "Description", "CreatedAt", "UpdatedAt", "CreatedBy", "UpdatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateTransaction :one
UPDATE "TransactionHistory"
SET "TransactionType" = $2, "Amount" = $3, "Description" = $4, "UpdatedAt" = $5, "UpdatedBy" = $6
WHERE "Id" = $1
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM "TransactionHistory" WHERE "Id" = $1;
