package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"v0/internal/db/sqlite"
	"v0/internal/delivery/http/handler"
	"v0/internal/middleware"
	"v0/internal/repository"
	"v0/internal/useCase"
)

func main() {

	db, err := sqlite.NewSQLiteDB()
	if err != nil {
		log.Fatalf("Error initializing SQLite DB: %v", err)
	}
	defer db.Close()

	// Repositories configurations
	jobRepository := repository.NewSQLiteJobRepository(db)
	externalRepository := repository.NewExternalAPIRepository()

	// Usecase configurations
	jobUseCase := useCase.NewJobUseCase(jobRepository, externalRepository)

	// Handler configurations
	jobHandler := handler.NewJobHandler(jobUseCase)

	// Create new router
	router := mux.NewRouter()
	http.Handle("/", middleware.JSONContentTypeMiddleware(router))

	// Register handlers
	jobsPath := "/jobs"
	router.HandleFunc(jobsPath, jobHandler.CreateJob).Methods(http.MethodPost)
	router.HandleFunc(jobsPath, jobHandler.GetJob).Methods(http.MethodGet)

	log.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
