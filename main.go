package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Bones1335/workout_api/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db       *database.Queries
	platform string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")

	const filepathRoot = "."
	const port = "8080"

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Error connecting to DB: %v", err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		db:       dbQueries,
		platform: platform,
	}

	mux := http.NewServeMux()

	fsHandler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot+"/front-end/.")))
	mux.Handle("/app/", fsHandler)

	// Admin endpoints
	mux.HandleFunc("GET /healthz", handlerReadiness)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	// User endpoints
	mux.HandleFunc("POST /api/users", apiCfg.handlerCreateUsers)
	mux.HandleFunc("GET /api/users/{user_id}", apiCfg.handlerGetUsers)
	mux.HandleFunc("PUT /api/users/{user_id}", apiCfg.handlerUpdateUsers)
	//TODO: mux.HandleFunc("DELETE /api/users/{user_id}", apiCfg.handlerDeleteUsers)

	// Exercise endpoints
	mux.HandleFunc("POST /api/exercises", apiCfg.handlerCreateExercises)
	mux.HandleFunc("GET /api/exercises", apiCfg.handlerGetExercises)
	mux.HandleFunc("GET /api/users/{user_id}/exercises", apiCfg.handlerGetUserExercises)
	mux.HandleFunc("GET /api/exercises/{exercise_id}", apiCfg.handlerGetSingleExercise)
	mux.HandleFunc("PUT /api/exercises/{exercise_id}", apiCfg.handlerUpdateExercises)
	mux.HandleFunc("DELETE /api/exercises/{exercise_id}", apiCfg.handlerDeleteExercises)

	// Workout Endpoints
	mux.HandleFunc("POST /api/workouts", apiCfg.handlerCreateWorkoutRoutines)
	mux.HandleFunc("GET /api/workouts", apiCfg.handleGetWorkoutRoutines)
	//TODO: mux.HandleFunc("GET /api/users/workouts/{workout_id}", apiCfg.handleGetWorkout)
	//TODO: mux.HandleFunc("PUT /api/workouts/{workout_id}", apiCfg.handleUpdateWorkout)
	//TODO: mux.HandleFun("DELETE /api/workouts/{workout_id}", apiCfg.handleDeleteWorkout)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())

}
