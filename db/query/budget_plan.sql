-- name: GetAllBudgetPlans :many
SELECT * FROM "BudgetPlan";

-- name: GetBudgetPlanById :one
SELECT * FROM "BudgetPlan" WHERE "Id" = $1;

-- name: CreateBudgetPlan :one
INSERT INTO "BudgetPlan" ("Id", "AccountId", "Month", "TotalIncome", "TotalExpenses", "SavingsGoal", "CreatedAt", "UpdatedAt", "CreatedBy", "UpdatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateBudgetPlan :one
UPDATE "BudgetPlan"
SET "Month" = $2, "TotalIncome" = $3, "TotalExpenses" = $4, "SavingsGoal" = $5, "UpdatedAt" = $6, "UpdatedBy" = $7
WHERE "Id" = $1
RETURNING *;

-- name: DeleteBudgetPlan :exec
DELETE FROM "BudgetPlan" WHERE "Id" = $1;
