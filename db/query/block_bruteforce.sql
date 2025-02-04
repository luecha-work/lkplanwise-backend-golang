-- name: CreateBlockBruteForce :one
INSERT INTO "BlockBruteForce" (
  "Email", "Count", "Status", 
  "LockedTime", "UnLockTime", "CreatedAt", "CreatedBy"
) 
VALUES (
  $1, $2, $3, 
  $4, $5, $6, $7
)
RETURNING *;

-- name: GetBlockBruteForceById :one
SELECT * 
FROM "BlockBruteForce" 
WHERE "Id" = $1 LIMIT 1;

-- name: GetBlockBruteForceByEmail :one
SELECT * 
FROM "BlockBruteForce" 
WHERE "Email" = $1 LIMIT 1;

-- name: UpdateBlockBruteForce :one
UPDATE "BlockBruteForce"
SET 
  "Email" = COALESCE(sqlc.narg(Email), "Email"),
  "Count" = COALESCE(sqlc.narg(Count), "Count"),
  "Status" = COALESCE(sqlc.narg(Status), "Status"),
  "LockedTime" = sqlc.arg(LockedTime),
  "UnLockTime" = sqlc.arg(UnLockTime),
  "UpdatedAt" = COALESCE(sqlc.narg(UpdatedAt), "UpdatedAt"),
  "UpdatedBy" = COALESCE(sqlc.narg(UpdatedBy), "UpdatedBy")
WHERE "Id" = sqlc.arg(Id)
RETURNING *;

-- name: DeleteBlockBruteForce :one
DELETE FROM "BlockBruteForce"
WHERE "Id" = $1
RETURNING *;
