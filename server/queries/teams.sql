-- name: getmany
SELECT *
FROM teams 
WHERE archived = false
AND id = ANY($1);

-- name: updatemany
UPDATE teams 
SET (school_id, name) = ($1, $2)
WHERE id = ANY($3)
RETURNING *;

-- name: deletemany
UPDATE teams 
SET archived = true
WHERE id = ANY($1)
RETURNING *;

-- name: reference
SELECT teams.*
FROM teams 
INNER JOIN "%s" on "%s".id = teams."%s"
WHERE teams.archived = false
AND "%s".id = $1;

-- name: list
SELECT *
FROM teams
WHERE archived=false;

-- name: create
INSERT INTO teams (school_id, name)
VALUES ($1,$2)
RETURNING *;

-- name: read
SELECT *
FROM teams
WHERE id=$1
AND archived=false;

-- name: update
UPDATE teams 
SET (school_id, name) = ($1, $2)
WHERE id = $3
AND archived=false
RETURNING *;

-- name: delete
UPDATE teams 
SET (archived, archived_on) = (true, NOW())
WHERE id = $1
RETURNING *;

