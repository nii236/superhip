-- name: getmany
SELECT *
FROM students 
WHERE archived = false
AND id = ANY($1);

-- name: updatemany
UPDATE students 
SET (school_id, name) = ($1, $2)
WHERE id = ANY($3)
RETURNING *;

-- name: deletemany
UPDATE students 
SET archived = true
WHERE id = ANY($1)
RETURNING *;

-- name: reference
SELECT students.*
FROM students 
INNER JOIN "%s" on "%s".id = students."%s"
WHERE students.archived = false
AND "%s".id = $1;

-- name: list
SELECT *
FROM students
WHERE archived=false
LIMIT $1
OFFSET $2;

-- name: create
INSERT INTO students (school_id, name)
VALUES ($1,$2)
RETURNING *;

-- name: read
SELECT *
FROM students
WHERE id=$1
AND archived=false;

-- name: update
UPDATE students 
SET (school_id, name) = ($1, $2)
WHERE id = $3
AND archived=false
RETURNING *;

-- name: delete
UPDATE students 
SET (archived, archived_on) = (true, NOW())
WHERE id = $1
RETURNING *;

