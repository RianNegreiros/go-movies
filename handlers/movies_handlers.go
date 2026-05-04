package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/riannegreiros/go-movies/data"
	"github.com/riannegreiros/go-movies/logger"
	"github.com/riannegreiros/go-movies/models"
)

type MoviesHandler struct {
	Storage data.MovieStorage
	Logger  *logger.Logger
}

// Utility functions
func (h *MoviesHandler) writeJSONResponse(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.Logger.Error("Failed to encode response", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (h *MoviesHandler) handleStorageError(w http.ResponseWriter, err error, context string) bool {
	if err != nil {
		if err == data.ErrMovieNotFound {
			http.Error(w, context, http.StatusNotFound)
			return true
		}
		h.Logger.Error(context, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return true
	}
	return false
}

func (h *MoviesHandler) parseID(w http.ResponseWriter, idStr string) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Error("Invalid ID format", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func (h *MoviesHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetTopMovies()

	if h.handleStorageError(w, err, "Failed to get movies") {
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.Logger.Info("Successfully served top movies")
	}
}

func (h *MoviesHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.Storage.GetRandomMovies()
	if h.handleStorageError(w, err, "Failed to get movies") {
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.Logger.Info("Successfully served random movies")
	}
}

func (h *MoviesHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	order := r.URL.Query().Get("order")
	genreStr := r.URL.Query().Get("genre")

	var genre *int
	if genreStr != "" {
		genreInt, ok := h.parseID(w, genreStr)
		if !ok {
			return
		}
		genre = &genreInt
	}

	var movies []models.Movie
	var err error
	if query != "" {
		movies, err = h.Storage.SearchMoviesByName(query, order, genre)
	}
	if h.handleStorageError(w, err, "Failed to get movies") {
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.Logger.Info("Successfully served movies")
	}
}

func (h *MoviesHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/movies/"):]
	id, ok := h.parseID(w, idStr)
	if !ok {
		return
	}

	movie, err := h.Storage.GetMovieByID(id)
	if h.handleStorageError(w, err, "Failed to get movie by ID") {
		return
	}
	if h.writeJSONResponse(w, movie) == nil {
		h.Logger.Info("Successfully served movie with ID: " + idStr)
	}
}

func (h *MoviesHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.Storage.GetAllGenres()
	if h.handleStorageError(w, err, "Failed to get genres") {
		return
	}
	if h.writeJSONResponse(w, genres) == nil {
		h.Logger.Info("Successfully served genres")
	}
}

func NewMoviesHandler(storage data.MovieStorage, log *logger.Logger) *MoviesHandler {
	return &MoviesHandler{
		Storage: storage,
		Logger:  log,
	}
}
