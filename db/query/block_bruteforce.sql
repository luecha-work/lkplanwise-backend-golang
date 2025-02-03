-- name: CreateBlockBruteForce :one
INSERT INTO "BlockBruteForce" (
  "UserName", "Count", "Status", 
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

-- name: GetBlockBruteForceByUsername :one
SELECT * 
FROM "BlockBruteForce" 
WHERE "UserName" = $1 LIMIT 1;

-- name: UpdateBlockBruteForce :one
UPDATE "BlockBruteForce"
SET 
  "UserName" = COALESCE(sqlc.narg(UserName), "UserName"),
  "Count" = COALESCE(sqlc.narg(Count), "Count"),
  "Status" = COALESCE(sqlc.narg(Status), "Status"),
  "LockedTime" = COALESCE(sqlc.narg(LockedTime), "LockedTime"),
  "UnLockTime" = COALESCE(sqlc.narg(UnLockTime), "UnLockTime"),
  "UpdatedAt" = COALESCE(sqlc.narg(UpdatedAt), "UpdatedAt"),
  "UpdatedBy" = COALESCE(sqlc.narg(UpdatedBy), "UpdatedBy")
WHERE "Id" = sqlc.arg(Id)
RETURNING *;

-- name: DeleteBlockBruteForce :one
DELETE FROM "BlockBruteForce"
WHERE "Id" = $1
RETURNING *;
