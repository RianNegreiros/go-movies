package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/riannegreiros/go-movies/data"
	"github.com/riannegreiros/go-movies/handlers"
	"github.com/riannegreiros/go-movies/logger"

	_ "github.com/lib/pq"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("go-movies.log")
	if err != nil {
		log.Fatalf("Could not create a log instance %v", err)
	}

	defer logInstance.Close()

	return logInstance
}

func main() {

	logInstance := initializeLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file available")
	}

	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATBASE_URL not set")
	}

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Close()

	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initalize movie repository: %v", err)
	}

	moviesHandler := handlers.MoviesHandler{}
	moviesHandler.Storage = movieRepo
	moviesHandler.Logger = logInstance

	http.HandleFunc("/api/movies/top", moviesHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", moviesHandler.GetRandomMovies)

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
		logInstance.Error("Could not start server: %v", err)
	}
}
