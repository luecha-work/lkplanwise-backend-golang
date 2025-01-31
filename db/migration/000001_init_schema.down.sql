-- Drop foreign key constraints first
ALTER TABLE public.acoount_roles DROP CONSTRAINT acoount_roles_fk;
ALTER TABLE public.acoount_roles DROP CONSTRAINT acoount_roles_fk_1;

-- Drop tables in reverse order to avoid dependency issues
DROP TABLE IF EXISTS public.acoount_roles;
DROP TABLE IF EXISTS public.roles;
DROP TABLE IF EXISTS public.accounts;