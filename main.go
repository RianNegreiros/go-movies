package main

import (
	"log"
	"net/http"

	"github.com/riannegreiros/go-movies/handlers"
	"github.com/riannegreiros/go-movies/logger"
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

	moviesHandler := handlers.MoviesHandler{}

	http.HandleFunc("/api/movies/top", moviesHandler.GetTopMovies)

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
		logInstance.Error("Could not start server: %v", err)
	}
}
