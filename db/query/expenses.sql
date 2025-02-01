-- name: GetAllExpenses :many
SELECT * FROM "Expense";

-- name: GetExpenseById :one
SELECT * FROM "Expense" WHERE "Id" = $1;

-- name: CreateExpense :one
INSERT INTO "Expense" ("Id", "AccountId", "Category", "Amount", "Date", "Description", "CreatedAt", "UpdatedAt", "CreatedBy", "UpdatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateExpense :one
UPDATE "Expense"
SET "Category" = $2, "Amount" = $3, "Date" = $4, "Description" = $5, "UpdatedAt" = $6, "UpdatedBy" = $7
WHERE "Id" = $1
RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM "Expense" WHERE "Id" = $1;
