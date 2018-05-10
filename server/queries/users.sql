-- name: getmany
SELECT *
FROM users 
WHERE archived = false
AND id = ANY($1);

-- name: updatemany
UPDATE users 
SET (first_name, last_name, email, password_hash, role) = ($1,$2,$3,$4,$5)
WHERE id = ANY($6)
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
SELECT 
users.*,
COALESCE(array_agg(schools.id), '{}'::UUID[])AS school_ids
FROM users
LEFT JOIN schools_users ON schools_users.user_id = users.id
LEFT JOIN schools ON schools_users.school_id = schools.id
WHERE users.archived=false
GROUP BY users.id




-- name: create
INSERT INTO users (first_name, last_name, email, password_hash, role)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: read
SELECT *
FROM users
WHERE id=$1
AND archived=false;

-- name: update
UPDATE users 
SET (first_name, last_name, email, password_hash, role) = ($1,$2,$3,$4,$5)
WHERE id = $6
AND archived=false
RETURNING *;

-- name: delete
UPDATE users 
SET (archived, archived_on) = (true, NOW())
WHERE id = $1
RETURNING *;

