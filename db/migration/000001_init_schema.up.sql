CREATE TABLE public.accounts (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	username varchar NOT NULL,
	firstname varchar NULL,
	lastname varchar NULL,
	email varchar NULL,
	password_hash varchar NULL,
	date_of_birth date NULL,
	created_at timestamptz NULL,
	created_by varchar NULL,
	updated_at timestamptz NULL,
	updated_by varchar NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_un UNIQUE (username)
);

CREATE TABLE public.roles (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	role_code varchar NOT NULL,
	role_name varchar NULL,
	created_at timestamptz NULL,
	created_by varchar NULL,
	updated_at timestamptz NULL,
	updated_by varchar NULL,
	CONSTRAINT roles_pk PRIMARY KEY (id)
);

CREATE TABLE public.account_roles (
	account_id uuid NOT NULL,
	role_id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	created_by varchar NULL,
	updated_by varchar NULL
);


-- public.account_roles foreign keys

ALTER TABLE public.account_roles ADD CONSTRAINT account_roles_fk FOREIGN KEY (account_id) REFERENCES public.accounts(id);
ALTER TABLE public.account_roles ADD CONSTRAINT account_roles_fk_1 FOREIGN KEY (role_id) REFERENCES public.roles(id);
