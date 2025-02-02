-- name: GetAllExpenses :many
SELECT * FROM "Expense";

-- name: GetExpenseById :one
SELECT * FROM "Expense" WHERE "Id" = $1;

-- name: CreateExpense :one
INSERT INTO "Expense" ("Id", "AccountId", "Category", "Amount", "Date", "Description", "CreatedAt", "CreatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateExpense :one
UPDATE "Expense"
SET 
  "Category" = COALESCE($2, "Category"),
  "Amount" = COALESCE($3, "Amount"),
  "Date" = COALESCE($4, "Date"),
  "Description" = COALESCE($5, "Description"),
  "UpdatedAt" = COALESCE($6, "UpdatedAt"),
  "UpdatedBy" = COALESCE($7, "UpdatedBy")
WHERE "Id" = $1
RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM "Expense" WHERE "Id" = $1;
