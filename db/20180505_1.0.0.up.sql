CREATE TABLE users (
	id serial NOT NULL,
	first_name text NOT NULL,
	last_name text NOT NULL,
	email text NOT NULL,
	password text NOT NULL,
	role text NOT NULL,
	session_token text NOT NULL,
	data jsonb NOT NULL,
	disabled boolean NOT NULL,
	disabled_on timestamp,
	archived boolean NOT NULL,
	archived_on timestamp,
	created_at timestamp NOT NULL,
	password_reset_token text NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_email_key UNIQUE (email)
);

