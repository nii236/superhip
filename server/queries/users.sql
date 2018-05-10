-- name: getmany
SELECT *
FROM users 
WHERE archived = false
AND id = ANY($1);

-- name: updatemany
UPDATE users 
SET (school_id, first_name, last_name, email, password_hash, role) = ($1,$2,$3,$4,$5,$6)
WHERE id = ANY($7)
RETURNING *;

-- name: deletemany
UPDATE users 
SET archived = true
WHERE id = ANY($1)
RETURNING *;

-- name: reference
SELECT users.*
FROM users 
INNER JOIN "%s" on "%s".id = users."%s"
WHERE users.archived = false
AND "%s".id = $1;

-- name: list
SELECT *
FROM users
WHERE archived=false;

-- name: create
INSERT INTO users (school_id, first_name, last_name, email, password_hash, role)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: read
SELECT *
FROM users
WHERE id=$1
AND archived=false;

-- name: update
UPDATE users 
SET (school_id, first_name, last_name, email, password_hash, role) = ($1,$2,$3,$4,$5,$6)
WHERE id = $7
AND archived=false
RETURNING *;

-- name: delete
UPDATE users 
SET (archived, archived_on) = (true, NOW())
WHERE id = $1
RETURNING *;

