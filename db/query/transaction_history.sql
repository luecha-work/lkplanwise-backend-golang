-- name: GetAllTransactions :many
SELECT * FROM "TransactionHistory";

-- name: GetTransactionById :one
SELECT * FROM "TransactionHistory" WHERE "Id" = $1 LIMIT 1;

-- name: CreateTransaction :one
INSERT INTO "TransactionHistory" ("Id", "AccountId", "TransactionType", "Amount", "Description", "CreatedAt", "CreatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateTransaction :one
UPDATE "TransactionHistory"
SET 
  "TransactionType" = COALESCE($2, "TransactionType"),
  "Amount" = COALESCE($3, "Amount"),
  "Description" = COALESCE($4, "Description"),
  "UpdatedAt" = COALESCE($5, "UpdatedAt"),
  "UpdatedBy" = COALESCE($6, "UpdatedBy")
WHERE "Id" = $1
RETURNING *;


-- name: DeleteTransaction :exec
DELETE FROM "TransactionHistory" WHERE "Id" = $1;
