-- name: getmany
SELECT *
FROM schools 
WHERE archived = false
AND id = ANY($1);

-- name: updatemany
UPDATE schools 
SET name = $1
WHERE id = ANY($2)
RETURNING *;

-- name: deletemany
UPDATE schools 
SET archived = true
WHERE id = ANY($1)
RETURNING *;

-- name: reference
SELECT schools.*
FROM schools 
INNER JOIN "%s" on "%s".id = schools."%s"
WHERE schools.archived = false
AND "%s".id = $1;

-- name: list
SELECT *
FROM schools
WHERE archived=false;

-- name: create
INSERT INTO schools (name)
VALUES ($1)
RETURNING *;

-- name: read
SELECT *
FROM schools
WHERE id=$1
AND archived=false;

-- name: update
UPDATE schools 
SET name = $1
WHERE id = $2
AND archived=false
RETURNING *;

-- name: delete
UPDATE schools 
SET (archived, archived_on) = (true, NOW())
WHERE id = $1
RETURNING *;

