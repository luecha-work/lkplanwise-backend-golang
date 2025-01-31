-- name: CreateRole :one
INSERT INTO public.roles (
  role_code, role_name, created_at, created_by
) VALUES (
  $1, $2, NOW(), $3
) RETURNING *;

-- name: GetRole :one
SELECT * FROM public.roles
WHERE id = $1 LIMIT 1;

-- name: GetRoleByCode :one
SELECT * FROM public.roles
WHERE role_code = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * FROM public.roles
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateRole :one
UPDATE public.roles
SET role_code = $2, role_name = $3, updated_at = NOW(), updated_by = $4
WHERE id = $1
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM public.roles
WHERE id = $1;