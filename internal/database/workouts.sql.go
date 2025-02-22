// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: workouts.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createWorkout = `-- name: CreateWorkout :one
INSERT INTO workouts (
    id,
    created_at,
    updated_at,
    name,
    description,
    total_duration,
    user_id
)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4
)
RETURNING id, created_at, updated_at, name, description, total_duration, user_id
`

type CreateWorkoutParams struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	TotalDuration int32     `json:"total_duration"`
	UserID        uuid.UUID `json:"user_id"`
}

func (q *Queries) CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (Workout, error) {
	row := q.db.QueryRowContext(ctx, createWorkout,
		arg.Name,
		arg.Description,
		arg.TotalDuration,
		arg.UserID,
	)
	var i Workout
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.TotalDuration,
		&i.UserID,
	)
	return i, err
}

const createWorkoutExercise = `-- name: CreateWorkoutExercise :one
INSERT INTO workouts_exercises (
    id,
    time_seconds,
    weight_kg,
    workout_id,
    exercise_id
)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4
)
RETURNING id, time_seconds, weight_kg, workout_id, exercise_id
`

type CreateWorkoutExerciseParams struct {
	TimeSeconds int32     `json:"time_seconds"`
	WeightKg    int32     `json:"weight_kg"`
	WorkoutID   uuid.UUID `json:"workout_id"`
	ExerciseID  uuid.UUID `json:"exercise_id"`
}

func (q *Queries) CreateWorkoutExercise(ctx context.Context, arg CreateWorkoutExerciseParams) (WorkoutsExercise, error) {
	row := q.db.QueryRowContext(ctx, createWorkoutExercise,
		arg.TimeSeconds,
		arg.WeightKg,
		arg.WorkoutID,
		arg.ExerciseID,
	)
	var i WorkoutsExercise
	err := row.Scan(
		&i.ID,
		&i.TimeSeconds,
		&i.WeightKg,
		&i.WorkoutID,
		&i.ExerciseID,
	)
	return i, err
}
