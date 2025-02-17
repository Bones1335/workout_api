// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: exercises.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createExercise = `-- name: CreateExercise :one
INSERT INTO exercises (
    id,
    created_at,
    updated_at,
    name,
    tool,
    user_id
)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING id, created_at, updated_at, name, tool, user_id
`

type CreateExerciseParams struct {
	Name   string    `json:"name"`
	Tool   string    `json:"tool"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) CreateExercise(ctx context.Context, arg CreateExerciseParams) (Exercise, error) {
	row := q.db.QueryRowContext(ctx, createExercise, arg.Name, arg.Tool, arg.UserID)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Tool,
		&i.UserID,
	)
	return i, err
}

const deleteExercise = `-- name: DeleteExercise :exec
DELETE FROM exercises
WHERE id = $1
`

func (q *Queries) DeleteExercise(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteExercise, id)
	return err
}

const getUserExercises = `-- name: GetUserExercises :many
SELECT id, created_at, updated_at, name, tool, user_id FROM exercises
WHERE user_id = $1
`

func (q *Queries) GetUserExercises(ctx context.Context, userID uuid.UUID) ([]Exercise, error) {
	rows, err := q.db.QueryContext(ctx, getUserExercises, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Exercise
	for rows.Next() {
		var i Exercise
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Tool,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExercise = `-- name: UpdateExercise :one
UPDATE exercises SET name = $2, tool = $3, updated_at = NOW()
WHERE id = $1
RETURNING id, created_at, updated_at, name, tool, user_id
`

type UpdateExerciseParams struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Tool string    `json:"tool"`
}

func (q *Queries) UpdateExercise(ctx context.Context, arg UpdateExerciseParams) (Exercise, error) {
	row := q.db.QueryRowContext(ctx, updateExercise, arg.ID, arg.Name, arg.Tool)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Tool,
		&i.UserID,
	)
	return i, err
}
