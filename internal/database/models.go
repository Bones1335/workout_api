// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Exercise struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Tool   string    `json:"tool"`
	UserID uuid.UUID `json:"user_id"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}

type WorkoutExercise struct {
	ID         uuid.UUID `json:"id"`
	WorkoutID  uuid.UUID `json:"workout_id"`
	ExerciseID uuid.UUID `json:"exercise_id"`
	Position   int32     `json:"position"`
}

type WorkoutRoutine struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	TotalDuration     int32     `json:"total_duration"`
	RoundsPerExercise int32     `json:"rounds_per_exercise"`
	RoundDuration     int32     `json:"round_duration"`
	RestDuration      int32     `json:"rest_duration"`
}
