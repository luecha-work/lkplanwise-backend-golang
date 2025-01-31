-- name: InsertAccount :exec
INSERT INTO public.accounts (
  username, firstname, lastname, email, password_hash, date_of_birth, created_at, created_by
) VALUES (
  $1, $2, $3, $4, $5, $6, NOW(), $7
);

-- name: UpdateAccount :exec
UPDATE public.accounts
SET username = $2, firstname = $3, lastname = $4, email = $5, password_hash = $6, date_of_birth = $7, updated_at = NOW(), updated_by = $8
WHERE id = $1;

-- name: GetAccountByID :one
SELECT id, username, firstname, lastname, email, date_of_birth, created_at, created_by, updated_at, updated_by
FROM public.accounts
WHERE id = $1
LIMIT 1;

-- name: DeleteAccountByID :exec
DELETE FROM public.accounts
WHERE id = $1;

-- name: ListAllAccounts :many
SELECT id, username, firstname, lastname, email, date_of_birth, created_at, created_by, updated_at, updated_by
FROM public.accounts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetAccountByUsername :one
SELECT id, username, firstname, lastname, email, date_of_birth, created_at, created_by, updated_at, updated_by
FROM public.accounts
WHERE username = $1
LIMIT 1;

-- name: GetAccountsByRole :many
SELECT a.* FROM public.accounts a
JOIN public.account_roles ar ON a.id = ar.account_id
WHERE ar.role_id = $1;
