-- name: GetAllGoals :many
SELECT * FROM "Goal";

-- name: GetGoalById :one
SELECT * FROM "Goal" WHERE "Id" = $1;

-- name: CreateGoal :one
INSERT INTO "Goal" ("Id", "AccountId", "GoalType", "TargetAmount", "CurrentAmount", "Deadline", "Progress", "CreatedAt", "UpdatedAt", "CreatedBy", "UpdatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: UpdateGoal :one
UPDATE "Goal"
SET "GoalType" = $2, "TargetAmount" = $3, "CurrentAmount" = $4, "Deadline" = $5, "Progress" = $6, "UpdatedAt" = $7, "UpdatedBy" = $8
WHERE "Id" = $1
RETURNING *;

-- name: DeleteGoal :exec
DELETE FROM "Goal" WHERE "Id" = $1;
