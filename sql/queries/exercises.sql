-- name: CreateExercise :one
INSERT INTO exercises (
    id,
    name,
    tool,
    user_id
)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetUserExercises :many
SELECT * FROM exercises
WHERE user_id = $1;

-- name: GetExercises :many
SELECT * FROM exercises;

-- name: GetSingleExercise :one
SELECT * FROM exercises
WHERE id = $1;

-- name: UpdateExercise :one
UPDATE exercises SET name = $2, tool = $3 
WHERE id = $1
RETURNING *;

-- name: DeleteExercise :exec
DELETE FROM exercises
WHERE id = $1;