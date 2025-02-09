-- name: GetAllAccounts :many
SELECT * FROM "Accounts";

-- name: GetAccountById :one
SELECT * FROM "Accounts" WHERE "Id" = $1 LIMIT 1;

-- name: GetAccountByUsername :one
SELECT * FROM "Accounts"
WHERE "UserName" = $1 LIMIT 1;

-- name: GetAccountByEmail :one
SELECT * FROM "Accounts"
WHERE "Email" = $1 LIMIT 1;

-- name: PagedAccounts :many
SELECT * FROM "Accounts"
ORDER BY "CreatedAt"
LIMIT $1
OFFSET $2;

-- name: CreateAccount :one
INSERT INTO "Accounts" ("Id", "UserName", "FirstName", "LastName", "Email", "PasswordHash", "DateOfBirth", "RoleId", "CreatedAt", "CreatedBy")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateAccount :one
UPDATE "Accounts"
SET 
  "UserName" = COALESCE(sqlc.narg(UserName), "UserName"),
  "FirstName" = COALESCE(sqlc.narg(FirstName), "FirstName"),
  "LastName" = COALESCE(sqlc.narg(LastName), "LastName"),
  "Email" = COALESCE(sqlc.narg(Email), "Email"),
  "PasswordHash" = COALESCE(sqlc.narg(PasswordHash), "PasswordHash"),
  "DateOfBirth" = COALESCE(sqlc.narg(DateOfBirth), "DateOfBirth"),
  "RoleId" = COALESCE(sqlc.narg(RoleId), "RoleId"),
  "UpdatedAt" = COALESCE(sqlc.narg(UpdatedAt), "UpdatedAt"),
  "UpdatedBy" = COALESCE(sqlc.narg(UpdatedBy), "UpdatedBy"),
  "IsMailVerified" = COALESCE(sqlc.narg(IsMailVerified), "IsMailVerified"),
  "IsLocked" = COALESCE(sqlc.narg(IsLocked), "IsLocked")
WHERE "Id" = sqlc.arg(Id)
RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM "Accounts" WHERE "Id" = $1;