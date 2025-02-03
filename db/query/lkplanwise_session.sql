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
WHERE "Id" = $1 LIMIT 1;

-- name: GetLKPlanWiseSessionForLogin :one
SELECT * 
FROM "LKPlanWiseSession" 
WHERE "AccountId" = $1 AND "LoginIp" = $2 LIMIT 1;

-- name: UpdateLKPlanWiseSession :one
UPDATE "LKPlanWiseSession"
SET 
  "AccountId" = COALESCE(sqlc.narg(AccountId), "AccountId"),
  "LoginAt" = COALESCE(sqlc.narg(LoginAt), "LoginAt"),
  "Platform" = COALESCE(sqlc.narg(Platform), "Platform"),
  "Os" = COALESCE(sqlc.narg(Os), "Os"),
  "Browser" = COALESCE(sqlc.narg(Browser), "Browser"),
  "LoginIp" = COALESCE(sqlc.narg(LoginIp), "LoginIp"),
  "IssuedTime" = COALESCE(sqlc.narg(IssuedTime), "IssuedTime"),
  "ExpirationTime" = COALESCE(sqlc.narg(ExpirationTime), "ExpirationTime"),
  "SessionStatus" = COALESCE(sqlc.narg(SessionStatus), "SessionStatus"),
  "Token" = COALESCE(sqlc.narg(Token), "Token"),
  "RefreshTokenAt" = COALESCE(sqlc.narg(RefreshTokenAt), "RefreshTokenAt"),
  "UpdatedAt" = COALESCE(sqlc.narg(UpdatedAt), "UpdatedAt"),
  "UpdatedBy" = COALESCE(sqlc.narg(UpdatedBy), "UpdatedBy")
WHERE "Id" = sqlc.arg(Id)
RETURNING *;


-- name: DeleteLKPlanWiseSession :one
DELETE FROM "LKPlanWiseSession"
WHERE "Id" = $1
RETURNING *;