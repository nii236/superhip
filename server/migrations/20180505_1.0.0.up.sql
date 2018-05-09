CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	first_name text NOT NULL,
	last_name text NOT NULL,
	email text NOT NULL,
	password_hash text NOT NULL,
	password_reset_token text NOT NULL DEFAULT 'none',
	role text NOT NULL DEFAULT 'teacher',
	metadata jsonb NOT NULL DEFAULT '{}',
	archived boolean NOT NULL DEFAULT false,
	archived_on timestamp,
	created_at timestamp NOT NULL DEFAULT NOW(),
	CONSTRAINT users_email_key UNIQUE (email)
);

CREATE TABLE schools (
	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL
);

CREATE TABLE teams (
	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	user_id UUID NOT NULL REFERENCES users(id),
	name TEXT NOT NULL,
	metadata jsonb NOT NULL DEFAULT '{}',
	archived boolean NOT NULL DEFAULT false,
	archived_on timestamp,
	created_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE students (
	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL,
	team_id UUID NOT NULL REFERENCES teams(id)
);

-- many to many relationship between identities and roles.
-- many to many relationship between roles and permissions.
-- roles can have a parent role (inheriting permissions).


-- CREATE TABLE roles (
-- 	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
-- 	parent_id UUID NOT NULL REFERENCES roles(id),
-- 	name TEXT NOT NULL
-- );

-- CREATE TABLE permissions (
-- 	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
-- 	name TEXT NOT NULL
-- );

-- CREATE TABLE roles_users (
-- 	role_id UUID NOT NULL REFERENCES roles(id),
-- 	user_id UUID NOT NULL REFERENCES users(id),
-- 	name TEXT NOT NULL
-- );


-- CREATE TABLE roles_permissions (
-- 	role_id UUID NOT NULL REFERENCES roles(id),
-- 	permission_id UUID NOT NULL REFERENCES permissions(id),
-- 	name TEXT NOT NULL
-- );
