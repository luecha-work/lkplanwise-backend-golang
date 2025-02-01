-- name: GetAllRoles :many
SELECT * FROM "Roles";

-- name: GetRoleById :one
SELECT * FROM "Roles" WHERE "Id" = $1;

-- name: CreateRole :one
INSERT INTO "Roles" ("Id", "RoleCode", "RoleName", "CreatedAt", "UpdatedAt", "CreatedBy", "UpdatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateRole :one
UPDATE "Roles"
SET "RoleCode" = $2, "RoleName" = $3, "UpdatedAt" = $4, "UpdatedBy" = $5
WHERE "Id" = $1
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM "Roles" WHERE "Id" = $1;
