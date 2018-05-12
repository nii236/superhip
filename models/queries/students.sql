-- name: getmany
SELECT students.*,
COALESCE(array_agg(DISTINCT teams.id), '{}'::UUID[]) AS team_ids
FROM students
LEFT JOIN teams_students ON teams_students.student_id = students.id
LEFT JOIN teams ON teams_students.team_id = teams.id
WHERE students.archived=false
AND students.id = ANY($1)
GROUP BY students.id

-- name: updatemany
UPDATE students 
SET school_id = $1, name = $2
FROM (
	SELECT students.*,
	COALESCE(array_agg(DISTINCT teams.id), '{}'::UUID[]) AS team_ids
	FROM students
	LEFT JOIN teams_students ON teams_students.student_id = students.id
	LEFT JOIN teams ON teams_students.team_id = teams.id
	WHERE students.archived=false
	AND students.id = ANY($3)
	GROUP BY students.id
) t
WHERE t.id = students.id
RETURNING t.*;

-- name: deletemany
UPDATE students 
SET archived=true
FROM (
	SELECT students.*,
	COALESCE(array_agg(DISTINCT teams.id), '{}'::UUID[]) AS team_ids
	FROM students
	LEFT JOIN teams_students ON teams_students.student_id = students.id
	LEFT JOIN teams ON teams_students.team_id = teams.id
	WHERE students.archived=false
	AND students.id = ANY($1)
	GROUP BY students.id
) t
WHERE t.id = students.id
RETURNING t.*;

-- name: reference
SELECT students.*
FROM students 
INNER JOIN "%s" on "%s".id = users."%s"
WHERE users.archived = false
AND "%s".id = $1;

-- name: list
SELECT 
t.*
FROM (
	SELECT students.*,
	COALESCE(array_agg(DISTINCT teams.id), '{}'::UUID[]) AS team_ids
	FROM students
	LEFT JOIN teams_students ON teams_students.student_id = students.id
	LEFT JOIN teams ON teams_students.team_id = teams.id
	WHERE students.archived=false
	GROUP BY students.id
) t
JOIN students ON students.id = t.id
WHERE students.id = t.id
LIMIT $1
OFFSET $2;


-- name: create
INSERT INTO students (school_id, name) VALUES($1,$2)
RETURNING *;

-- name: read
SELECT 
t.*
FROM (
	SELECT students.*,
	COALESCE(array_agg(DISTINCT teams.id), '{}'::UUID[]) AS team_ids
	FROM students
	LEFT JOIN teams_students ON teams_students.student_id = students.id
	LEFT JOIN teams ON teams_students.team_id = teams.id
	WHERE students.archived=false
	AND students.id = $1
	GROUP BY students.id
) t
JOIN students ON students.id = t.id;

-- name: update
UPDATE students 
SET school_id = $1, name = $2
FROM (
	SELECT students.*,
	COALESCE(array_agg(DISTINCT teams.id), '{}'::UUID[]) AS team_ids
	FROM students
	LEFT JOIN teams_students ON teams_students.student_id = students.id
	LEFT JOIN teams ON teams_students.team_id = teams.id
	WHERE students.archived=false
	AND students.id = $3
	GROUP BY students.id
) t
WHERE t.id = students.id
RETURNING t.*;

-- name: delete
UPDATE students 
SET archived=true
FROM (
	SELECT students.*,
	COALESCE(array_agg(DISTINCT teams.id), '{}'::UUID[]) AS team_ids
	FROM students
	LEFT JOIN teams_students ON teams_students.student_id = students.id
	LEFT JOIN teams ON teams_students.team_id = teams.id
	WHERE students.archived=false
	AND students.id = $1
	GROUP BY students.id
) t
WHERE t.id = students.id
RETURNING t.*;
