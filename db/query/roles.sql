-- name: GetAllRoles :many
SELECT * FROM "Roles";

-- name: GetRoleById :one
SELECT * FROM "Roles" WHERE "Id" = $1 LIMIT 1;

-- name: CreateRole :one
INSERT INTO "Roles" ("Id", "RoleCode", "RoleName", "CreatedAt", "CreatedBy")
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateRole :one
UPDATE "Roles"
SET 
  "RoleCode" = COALESCE($2, "RoleCode"),
  "RoleName" = COALESCE($3, "RoleName"),
  "UpdatedAt" = COALESCE($4, "UpdatedAt"),
  "UpdatedBy" = COALESCE($5, "UpdatedBy")
WHERE "Id" = $1
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM "Roles" WHERE "Id" = $1;
