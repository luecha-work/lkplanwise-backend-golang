-- name: InsertAccountRole :exec
INSERT INTO public.account_roles (
  account_id, role_id, created_at, created_by
) VALUES (
  $1, $2, NOW(), $3
);

-- name: RemoveRoleFromAccount :exec
DELETE FROM public.account_roles
WHERE account_id = $1 AND role_id = $2;

-- name: GetAllAccountRoles :many
SELECT * FROM public.account_roles;

-- name: GetAccountRolesByAccountId :many
SELECT * FROM public.account_roles WHERE account_id = $1;

-- name: GetAccountRolesByRoleId :many
SELECT * FROM public.account_roles WHERE role_id = $1;

-- name: UpdateAccountsRoles :exec
UPDATE public.account_roles
SET 
    updated_at = NOW(),
    updated_by = $3
WHERE account_id = $1 AND role_id = $2;