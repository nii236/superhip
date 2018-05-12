-- name: getmany
SELECT users.*,
COALESCE(array_agg(DISTINCT schools.id), '{}'::UUID[]) AS school_ids,
COALESCE(array_agg(DISTINCT roles.id), '{}'::UUID[]) AS role_ids
FROM users
LEFT JOIN schools_users ON schools_users.user_id = users.id
LEFT JOIN schools ON schools_users.school_id = schools.id
LEFT JOIN roles_users ON roles_users.user_id = users.id
LEFT JOIN roles ON roles_users.role_id = roles.id
WHERE users.archived=false
AND users.id = ANY($1)
GROUP BY users.id

-- name: updatemany
UPDATE users 
SET (first_name, last_name, email, password_hash) = ($1,$2,$3,$4)
FROM (
	SELECT users.*,
	COALESCE(array_agg(DISTINCT schools.id), '{}'::UUID[]) AS school_ids,
	COALESCE(array_agg(DISTINCT roles.id), '{}'::UUID[]) AS role_ids
	FROM users
	LEFT JOIN schools_users ON schools_users.user_id = users.id
	LEFT JOIN schools ON schools_users.school_id = schools.id
	LEFT JOIN roles_users ON roles_users.user_id = users.id
	LEFT JOIN roles ON roles_users.role_id = roles.id
	WHERE users.archived=false
	AND users.id = ANY($5)
	GROUP BY users.id
) u
WHERE u.id = users.id
RETURNING u.*;

-- name: deletemany
UPDATE users 
SET archived=true
FROM (
	SELECT users.*,
	COALESCE(array_agg(DISTINCT schools.id), '{}'::UUID[]) AS school_ids,
	COALESCE(array_agg(DISTINCT roles.id), '{}'::UUID[]) AS role_ids
	FROM users
	LEFT JOIN schools_users ON schools_users.user_id = users.id
	LEFT JOIN schools ON schools_users.school_id = schools.id
	LEFT JOIN roles_users ON roles_users.user_id = users.id
	LEFT JOIN roles ON roles_users.role_id = roles.id
	WHERE users.archived=false
	AND users.id = ANY($1)
	GROUP BY users.id
) u
WHERE u.id = users.id
RETURNING u.*;

-- name: reference
SELECT users.*
FROM users 
INNER JOIN "%s" on "%s".id = users."%s"
WHERE users.archived = false
AND "%s".id = $1;

-- name: list
SELECT 
u.*
FROM (
	SELECT users.*,
	COALESCE(array_agg(DISTINCT schools.id), '{}'::UUID[]) AS school_ids,
	COALESCE(array_agg(DISTINCT roles.id), '{}'::UUID[]) AS role_ids
	FROM users
	LEFT JOIN schools_users ON schools_users.user_id = users.id
	LEFT JOIN schools ON schools_users.school_id = schools.id
	LEFT JOIN roles_users ON roles_users.user_id = users.id
	LEFT JOIN roles ON roles_users.role_id = roles.id
	WHERE users.archived=false
	GROUP BY users.id
) u
JOIN users ON u.id = users.id
WHERE users.id = u.id;


-- name: create
INSERT INTO users (first_name, last_name, email, password_hash)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: read
SELECT 
u.*
FROM (
	SELECT users.*,
	COALESCE(array_agg(DISTINCT schools.id), '{}'::UUID[]) AS school_ids,
	COALESCE(array_agg(DISTINCT roles.id), '{}'::UUID[]) AS role_ids
	FROM users
	LEFT JOIN schools_users ON schools_users.user_id = users.id
	LEFT JOIN schools ON schools_users.school_id = schools.id
	LEFT JOIN roles_users ON roles_users.user_id = users.id
	LEFT JOIN roles ON roles_users.role_id = roles.id
	WHERE users.archived=false
	AND users.id = $1
	GROUP BY users.id
) u
JOIN users ON u.id = users.id;

-- name: update
UPDATE users 
SET (first_name, last_name, email, password_hash) = ($1,$2,$3,$4)
FROM (
	SELECT users.*,
	COALESCE(array_agg(DISTINCT schools.id), '{}'::UUID[]) AS school_ids,
	COALESCE(array_agg(DISTINCT roles.id), '{}'::UUID[]) AS role_ids
	FROM users
	LEFT JOIN schools_users ON schools_users.user_id = users.id
	LEFT JOIN schools ON schools_users.school_id = schools.id
	LEFT JOIN roles_users ON roles_users.user_id = users.id
	LEFT JOIN roles ON roles_users.role_id = roles.id
	WHERE users.archived=false
	AND users.id = $5
	GROUP BY users.id
) u
WHERE u.id = users.id
RETURNING u.*;

-- name: delete
UPDATE users 
SET archived=true
FROM (
	SELECT users.*,
	COALESCE(array_agg(DISTINCT schools.id), '{}'::UUID[]) AS school_ids,
	COALESCE(array_agg(DISTINCT roles.id), '{}'::UUID[]) AS role_ids
	FROM users
	LEFT JOIN schools_users ON schools_users.user_id = users.id
	LEFT JOIN schools ON schools_users.school_id = schools.id
	LEFT JOIN roles_users ON roles_users.user_id = users.id
	LEFT JOIN roles ON roles_users.role_id = roles.id
	WHERE users.archived=false
	AND users.id = $1
	GROUP BY users.id
) u
WHERE u.id = users.id
RETURNING u.*;
