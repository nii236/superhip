-- name: getmany
SELECT teams.*,
COALESCE(array_agg(DISTINCT users.id), '{}'::UUID[]) AS user_ids,
COALESCE(array_agg(DISTINCT students.id), '{}'::UUID[]) AS student_ids
FROM teams
LEFT JOIN teams_users ON teams_users.team_id = teams.id
LEFT JOIN users ON teams_users.user_id = users.id
LEFT JOIN teams_students ON teams_students.team_id = teams.id
LEFT JOIN students ON teams_students.student_id = students.id
WHERE teams.archived=false
AND teams.id = ANY($1)
GROUP BY teams.id

-- name: updatemany
UPDATE teams 
SET school_id = $1, name = $2
FROM (
	SELECT teams.*,
	COALESCE(array_agg(DISTINCT users.id), '{}'::UUID[]) AS user_ids,
	COALESCE(array_agg(DISTINCT students.id), '{}'::UUID[]) AS student_ids
	FROM teams
	LEFT JOIN teams_users ON teams_users.team_id = teams.id
	LEFT JOIN users ON teams_users.user_id = users.id
	LEFT JOIN teams_students ON teams_students.team_id = teams.id
	LEFT JOIN students ON teams_students.student_id = students.id
	WHERE teams.archived=false
	AND teams.id = ANY($3)
	GROUP BY teams.id
) t
WHERE t.id = teams.id
RETURNING t.*;

-- name: deletemany
UPDATE teams 
SET archived=true
FROM (
	SELECT teams.*,
	COALESCE(array_agg(DISTINCT users.id), '{}'::UUID[]) AS user_ids,
	COALESCE(array_agg(DISTINCT students.id), '{}'::UUID[]) AS student_ids
	FROM teams
	LEFT JOIN teams_users ON teams_users.team_id = teams.id
	LEFT JOIN users ON teams_users.user_id = users.id
	LEFT JOIN teams_students ON teams_students.team_id = teams.id
	LEFT JOIN students ON teams_students.student_id = students.id
	WHERE teams.archived=false
	AND teams.id = ANY($1)
	GROUP BY teams.id
) t
WHERE t.id = teams.id
RETURNING t.*;

-- name: reference
SELECT teams.*
FROM teams 
INNER JOIN "%s" on "%s".id = users."%s"
WHERE users.archived = false
AND "%s".id = $1;

-- name: list
SELECT t.*
FROM (
	SELECT teams.*,
	COALESCE(array_agg(DISTINCT users.id), '{}'::UUID[]) AS user_ids,
	COALESCE(array_agg(DISTINCT students.id), '{}'::UUID[]) AS student_ids
	FROM teams
	LEFT JOIN teams_users ON teams_users.team_id = teams.id
	LEFT JOIN users ON teams_users.user_id = users.id
	LEFT JOIN teams_students ON teams_students.team_id = teams.id
	LEFT JOIN students ON teams_students.student_id = students.id
	WHERE teams.archived=false
	GROUP BY teams.id
) t
JOIN teams ON teams.id = t.id
WHERE teams.id = t.id
LIMIT $1
OFFSET $2;


-- name: create
INSERT INTO teams (school_id, name) VALUES ($1, $2)
RETURNING *;

-- name: read
SELECT 
t.*
FROM (
	SELECT teams.*,
	COALESCE(array_agg(DISTINCT users.id), '{}'::UUID[]) AS user_ids,
	COALESCE(array_agg(DISTINCT students.id), '{}'::UUID[]) AS student_ids
	FROM teams
	LEFT JOIN teams_users ON teams_users.team_id = teams.id
	LEFT JOIN users ON teams_users.user_id = users.id
	LEFT JOIN teams_students ON teams_students.team_id = teams.id
	LEFT JOIN students ON teams_students.student_id = students.id
	WHERE teams.archived=false
	AND teams.id = $1
	GROUP BY teams.id
) t
JOIN teams ON teams.id = t.id;

-- name: update
UPDATE teams 
SET school_id = $1, name = $2
FROM (
	SELECT teams.*,
	COALESCE(array_agg(DISTINCT users.id), '{}'::UUID[]) AS user_ids,
	COALESCE(array_agg(DISTINCT students.id), '{}'::UUID[]) AS student_ids
	FROM teams
	LEFT JOIN teams_users ON teams_users.team_id = teams.id
	LEFT JOIN users ON teams_users.user_id = users.id
	LEFT JOIN teams_students ON teams_students.team_id = teams.id
	LEFT JOIN students ON teams_students.student_id = students.id
	WHERE teams.archived=false
	AND teams.id = $3
	GROUP BY teams.id
) t
WHERE t.id = teams.id
RETURNING t.*;

-- name: delete
UPDATE teams 
SET archived=true
FROM (
	SELECT teams.*,
	COALESCE(array_agg(DISTINCT users.id), '{}'::UUID[]) AS user_ids,
	COALESCE(array_agg(DISTINCT students.id), '{}'::UUID[]) AS student_ids
	FROM teams
	LEFT JOIN teams_users ON teams_users.team_id = teams.id
	LEFT JOIN users ON teams_users.user_id = users.id
	LEFT JOIN teams_students ON teams_students.team_id = teams.id
	LEFT JOIN students ON teams_students.student_id = students.id
	WHERE teams.archived=false
	AND teams.id = $1
	GROUP BY teams.id
) t
WHERE t.id = teams.id
RETURNING t.*;
