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
WHERE "Id" = $1;

-- name: UpdateBlockBruteForce :one
UPDATE "BlockBruteForce"
SET 
  "Email" = COALESCE($2, "Email"),
  "Count" = COALESCE($3, "Count"),
  "Status" = COALESCE($4, "Status"),
  "LockedTime" = COALESCE($5, "LockedTime"),
  "UnLockTime" = COALESCE($6, "UnLockTime"),
  "UpdatedAt" = COALESCE($7, "UpdatedAt"),
  "UpdatedBy" = COALESCE($8, "UpdatedBy")
WHERE "Id" = $1
RETURNING *;

-- name: DeleteBlockBruteForce :one
DELETE FROM "BlockBruteForce"
WHERE "Id" = $1
RETURNING *;
