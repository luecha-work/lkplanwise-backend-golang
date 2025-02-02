-- name: CreateLKPlanWiseSession :one
INSERT INTO "LKPlanWiseSession" (
  "AccountId", "LoginAt", "Platform", "Os", "Browser", 
  "LoginIp", "IssuedTime", "ExpirationTime", "SessionStatus", 
  "Token", "RefreshTokenAt", "CreatedAt", "CreatedBy"
) 
VALUES (
  $1, $2, $3, $4, $5, 
  $6, $7, $8, $9, 
  $10, $11, $12, $13
) 
RETURNING *;

-- name: GetLKPlanWiseSessionById :one
SELECT * 
FROM "LKPlanWiseSession" 
WHERE "Id" = $1;

-- name: GetLKPlanWiseSessionForLogin :one
SELECT * 
FROM "LKPlanWiseSession" 
WHERE "AccountId" = $1 AND "LoginIp" = $2;

-- name: UpdateLKPlanWiseSession :one
UPDATE "LKPlanWiseSession"
SET 
  "AccountId" = COALESCE($2, "AccountId"),
  "LoginAt" = COALESCE($3, "LoginAt"),
  "Platform" = COALESCE($4, "Platform"),
  "Os" = COALESCE($5, "Os"),
  "Browser" = COALESCE($6, "Browser"),
  "LoginIp" = COALESCE($7, "LoginIp"),
  "IssuedTime" = COALESCE($8, "IssuedTime"),
  "ExpirationTime" = COALESCE($9, "ExpirationTime"),
  "SessionStatus" = COALESCE($10, "SessionStatus"),
  "Token" = COALESCE($11, "Token"),
  "RefreshTokenAt" = COALESCE($12, "RefreshTokenAt"),
  "UpdatedAt" = COALESCE($13, "UpdatedAt"),
  "UpdatedBy" = COALESCE($14, "UpdatedBy")
WHERE "Id" = $1
RETURNING *;

-- name: DeleteLKPlanWiseSession :one
DELETE FROM "LKPlanWiseSession"
WHERE "Id" = $1
RETURNING *;