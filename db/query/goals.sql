-- name: GetAllGoals :many
SELECT * FROM "Goal";

-- name: GetGoalById :one
SELECT * FROM "Goal" WHERE "Id" = $1;

-- name: CreateGoal :one
INSERT INTO "Goal" ("Id", "AccountId", "GoalType", "TargetAmount", "CurrentAmount", "Deadline", "Progress", "CreatedAt", "CreatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateGoal :one
UPDATE "Goal"
SET 
  "GoalType" = COALESCE($2, "GoalType"),
  "TargetAmount" = COALESCE($3, "TargetAmount"),
  "CurrentAmount" = COALESCE($4, "CurrentAmount"),
  "Deadline" = COALESCE($5, "Deadline"),
  "Progress" = COALESCE($6, "Progress"),
  "UpdatedAt" = COALESCE($7, "UpdatedAt"),
  "UpdatedBy" = COALESCE($8, "UpdatedBy")
WHERE "Id" = $1
RETURNING *;

-- name: DeleteGoal :exec
DELETE FROM "Goal" WHERE "Id" = $1;
