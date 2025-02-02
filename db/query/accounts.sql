-- name: GetAllAccounts :many
SELECT * FROM "Accounts";

-- name: GetAccountById :one
SELECT * FROM "Accounts" WHERE "Id" = $1;

-- name: GetAccountByUsername :one
SELECT * FROM "Accounts"
WHERE "UserName" = $1 LIMIT 1;

-- name: CreateAccount :one
INSERT INTO "Accounts" ("Id", "UserName", "FirstName", "LastName", "Email", "PasswordHash", "DateOfBirth", "RoleId", "CreatedAt", "CreatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateAccount :one
UPDATE "Accounts"
SET 
  "UserName" = COALESCE($2, "UserName"),
  "FirstName" = COALESCE($3, "FirstName"),
  "LastName" = COALESCE($4, "LastName"),
  "Email" = COALESCE($5, "Email"),
  "PasswordHash" = COALESCE($6, "PasswordHash"),
  "DateOfBirth" = COALESCE($7, "DateOfBirth"),
  "RoleId" = COALESCE($8, "RoleId"),
  "UpdatedAt" = COALESCE($9, "UpdatedAt"),
  "UpdatedBy" = COALESCE($10, "UpdatedBy")
WHERE "Id" = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM "Accounts" WHERE "Id" = $1;