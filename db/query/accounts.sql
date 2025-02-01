-- name: GetAllAccounts :many
SELECT * FROM "Accounts";

-- name: GetAccountById :one
SELECT * FROM "Accounts" WHERE "Id" = $1;

-- name: CreateAccount :one
INSERT INTO "Accounts" ("Id", "UserName", "FirstName", "LastName", "Email", "PasswordHash", "DateOfBirth", "RoleId", "CreatedAt", "UpdatedAt", "CreatedBy", "UpdatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: UpdateAccount :one
UPDATE "Accounts"
SET "UserName" = $2, "FirstName" = $3, "LastName" = $4, "Email" = $5, "PasswordHash" = $6, "DateOfBirth" = $7, "RoleId" = $8, "UpdatedAt" = $9, "UpdatedBy" = $10
WHERE "Id" = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM "Accounts" WHERE "Id" = $1;