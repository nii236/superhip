CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
	id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	first_name text NOT NULL,
	last_name text NOT NULL,
	email text NOT NULL,
	password_hash text NOT NULL,
	password_reset_token text NOT NULL DEFAULT 'none',
	role text NOT NULL DEFAULT 'teacher',
	data jsonb NOT NULL DEFAULT '{}',
	archived boolean NOT NULL DEFAULT false,
	archived_on timestamp,
	created_at timestamp NOT NULL DEFAULT NOW(),
	CONSTRAINT users_email_key UNIQUE (email)
);

CREATE TABLE schools (
	id UUID PRIMARY KEY NOT NULL,
	name TEXT NOT NULL
);


CREATE TABLE teams (
	id UUID PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	user_id UUID NOT NULL REFERENCES users(id)
);

CREATE TABLE students (
	id UUID PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	team_id UUID NOT NULL REFERENCES teams(id)
);
