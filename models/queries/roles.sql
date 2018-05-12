-- name: getmany
SELECT *
FROM roles 
WHERE archived = false
AND id = ANY($1);

-- name: updatemany
UPDATE roles 
SET name = $1
WHERE id = ANY($2)
RETURNING *;

-- name: deletemany
UPDATE roles 
SET archived = true
WHERE id = ANY($1)
RETURNING *;

-- name: reference
SELECT roles.*
FROM roles 
INNER JOIN "%s" on "%s".id = roles."%s"
WHERE roles.archived = false
AND "%s".id = $1;

-- name: list
SELECT *
FROM roles
WHERE archived=false
LIMIT $1
OFFSET $2;

-- name: create
INSERT INTO roles (name)
VALUES ($1)
RETURNING *;

-- name: read
SELECT *
FROM roles
WHERE id=$1
AND archived=false;

-- name: update
UPDATE roles 
SET name = $1
WHERE id = $2
AND archived=false
RETURNING *;

-- name: delete
UPDATE roles 
SET (archived, archived_on) = (true, NOW())
WHERE id = $1
RETURNING *;

