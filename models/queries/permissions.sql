-- name: getmany
SELECT *
FROM permissions 
WHERE archived = false
AND id = ANY($1);

-- name: updatemany
UPDATE permissions 
SET name = $1
WHERE id = ANY($2)
RETURNING *;

-- name: deletemany
UPDATE permissions 
SET archived = true
WHERE id = ANY($1)
RETURNING *;

-- name: reference
SELECT permissions.*
FROM permissions 
INNER JOIN "%s" on "%s".id = permissions."%s"
WHERE permissions.archived = false
AND "%s".id = $1;

-- name: list
SELECT *
FROM permissions
WHERE archived=false
LIMIT $1
OFFSET $2;

-- name: create
INSERT INTO permissions (name)
VALUES ($1)
RETURNING *;

-- name: read
SELECT *
FROM permissions
WHERE id=$1
AND archived=false;

-- name: update
UPDATE permissions 
SET name = $1
WHERE id = $2
AND archived=false
RETURNING *;

-- name: delete
UPDATE permissions 
SET (archived, archived_on) = (true, NOW())
WHERE id = $1
RETURNING *;

