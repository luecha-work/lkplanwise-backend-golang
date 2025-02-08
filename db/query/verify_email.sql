-- name: CreateVerifyEmail :one
INSERT INTO "VerifyEmails" (
    "UserName",
    "Email",
    "SecretCode"
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateVerifyEmail :one
UPDATE "VerifyEmails"
SET
    "IsUsed" = TRUE
WHERE
    "Id" = @Id
    AND "SecretCode" = @SecretCode
    AND "IsUsed" = FALSE
    AND "ExpiredAt" > now()
RETURNING *;