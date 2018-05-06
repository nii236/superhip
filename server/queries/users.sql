-- name: getmanyreference
SELECT *
FROM $1
WHERE archived=false
AND user_id = $2

-- name: all
SELECT id, first_name, last_name, email, role
FROM users
WHERE archived=false

-- name: create
INSERT INTO users (first_name, last_name, email, password_hash, role)
VALUES (:first_name, :last_name, :email, :password_hash, :role)
RETURNING id, first_name, last_name, email, password_hash, role

-- name: archive
UPDATE users SET (archived, archived_on)
= (true, NOW())
WHERE id = $1
RETURNING id, first_name, last_name, email, password_hash, role

-- name: get
SELECT id, first_name, last_name, email, role
FROM users
WHERE id=$1
AND archived=false

-- name: update
UPDATE users SET (first_name, last_name, email, password_hash, role)
= (:first_name, :last_name, :email, :password_hash, :role)
WHERE id = :id
AND archived=false
RETURNING id, first_name, last_name, email, password_hash, role
