-- name: GetAllBudgetPlans :many
SELECT * FROM "BudgetPlan";

-- name: GetBudgetPlanById :one
SELECT * FROM "BudgetPlan" WHERE "Id" = $1;

-- name: CreateBudgetPlan :one
INSERT INTO "BudgetPlan" ("Id", "AccountId", "Month", "TotalIncome", "TotalExpenses", "SavingsGoal", "CreatedAt", "CreatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateBudgetPlan :one
UPDATE "BudgetPlan"
SET 
  "Month" = COALESCE($2, "Month"),
  "TotalIncome" = COALESCE($3, "TotalIncome"),
  "TotalExpenses" = COALESCE($4, "TotalExpenses"),
  "SavingsGoal" = COALESCE($5, "SavingsGoal"),
  "UpdatedAt" = COALESCE($6, "UpdatedAt"),
  "UpdatedBy" = COALESCE($7, "UpdatedBy")
WHERE "Id" = $1
RETURNING *;

-- name: DeleteBudgetPlan :exec
DELETE FROM "BudgetPlan" WHERE "Id" = $1;
