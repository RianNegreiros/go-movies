package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/riannegreiros/go-movies/data"
	"github.com/riannegreiros/go-movies/handlers"
	"github.com/riannegreiros/go-movies/logger"

	_ "github.com/lib/pq"
)

func routes(accountHandler *handlers.AccountHandler, moviesHandler *handlers.MoviesHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/account/register", accountHandler.Register)
	mux.HandleFunc("POST /api/account/authenticate", accountHandler.Authenticate)

	mux.Handle("/api/account/favorites",
		accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.GetFavorites)))
	mux.Handle("/api/account/watchlist",
		accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.GetWatchlist)))
	mux.Handle("/api/account/save-to-collection",
		accountHandler.AuthMiddleware(http.HandlerFunc(accountHandler.SaveToCollection)))

	mux.HandleFunc("GET /api/movies/top", moviesHandler.GetTopMovies)
	mux.HandleFunc("GET /api/movies/random", moviesHandler.GetRandomMovies)
	mux.HandleFunc("GET /api/movies/search", moviesHandler.SearchMovies)
	mux.HandleFunc("GET /api/movies/{id}", moviesHandler.GetMovie)
	mux.HandleFunc("GET /api/genres", moviesHandler.GetGenres)

	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat("./public" + r.URL.Path); os.IsNotExist(err) {
			http.ServeFile(w, r, "./public/index.html")
			return
		}
		fs.ServeHTTP(w, r)
	}))

	return mux
}

func main() {
	logInstance, err := logger.NewLogger("go-movies.log")
	if err != nil {
		log.Fatalf("Could not create a log instance %v", err)
	}
	defer logInstance.Close()

	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal("No .env file available")
	}

	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	accountRepo, err := data.NewAccountRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize account repository: %v", err)
	}

	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize movie repository: %v", err)
	}

	mux := routes(
		handlers.NewAccountHandler(accountRepo, logInstance),
		handlers.NewMoviesHandler(movieRepo, logInstance),
	)

	fmt.Println("Serving on :8080")
	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
