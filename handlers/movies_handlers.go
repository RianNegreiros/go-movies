package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/riannegreiros/go-movies/models"
)

type MoviesHandler struct {
}

func (h *MoviesHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	overview1 := "A team of scientists embark on an expedition to the ocean's depths."
	overview2 := "In a post-apocalyptic world, a lone warrior fights for survival."
	overview3 := "A brilliant detective comes out of retirement to solve one last case."
	score1 := float32(8.5)
	score2 := float32(7.2)
	score3 := float32(9.0)
	pop1 := float32(120.5)
	pop2 := float32(85.3)
	pop3 := float32(200.1)
	lang1 := "en"
	lang2 := "en"
	lang3 := "fr"
	poster1 := "https://image.tmdb.org/t/p/w500/poster1.jpg"
	poster2 := "https://image.tmdb.org/t/p/w500/poster2.jpg"
	poster3 := "https://image.tmdb.org/t/p/w500/poster3.jpg"
	trailer1 := "https://youtube.com/watch?v=abc123"
	trailer2 := "https://youtube.com/watch?v=def456"
	trailer3 := "https://youtube.com/watch?v=ghi789"
	img1 := "https://image.tmdb.org/t/p/w200/actor1.jpg"
	img2 := "https://image.tmdb.org/t/p/w200/actor2.jpg"

	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     550,
			Title:       "Ocean's Mystery",
			Tagline:     "Dive deep into the unknown",
			ReleaseYear: 2023,
			Genres:      []models.Genre{{ID: 1, Name: "Adventure"}, {ID: 2, Name: "Sci-Fi"}},
			Overview:    &overview1,
			Score:       &score1,
			Popularity:  &pop1,
			Keywords:    []string{"ocean", "mystery", "science", "expedition"},
			Language:    &lang1,
			PosterURL:   &poster1,
			TrailerURL:  &trailer1,
			Casting: []models.Actor{
				{ID: 1, FistName: "Chris", LastName: "Evans", ImageURL: &img1},
				{ID: 2, FistName: "Scarlett", LastName: "Johansson", ImageURL: &img2},
			},
		},
		{
			ID:          2,
			TMDB_ID:     551,
			Title:       "Last Stand",
			Tagline:     "The end is just the beginning",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 3, Name: "Action"}, {ID: 4, Name: "Drama"}},
			Overview:    &overview2,
			Score:       &score2,
			Popularity:  &pop2,
			Keywords:    []string{"post-apocalyptic", "survival", "warrior", "action"},
			Language:    &lang2,
			PosterURL:   &poster2,
			TrailerURL:  &trailer2,
			Casting: []models.Actor{
				{ID: 3, FistName: "Tom", LastName: "Hardy"},
				{ID: 4, FistName: "Charlize", LastName: "Theron", ImageURL: &img2},
			},
		},
		{
			ID:          3,
			TMDB_ID:     552,
			Title:       "The Forgotten Case",
			Tagline:     "Some mysteries are better left unsolved",
			ReleaseYear: 2024,
			Genres:      []models.Genre{{ID: 5, Name: "Crime"}, {ID: 6, Name: "Thriller"}},
			Overview:    &overview3,
			Score:       &score3,
			Popularity:  &pop3,
			Keywords:    []string{"detective", "retirement", "case", "mystery", "thriller"},
			Language:    &lang3,
			PosterURL:   &poster3,
			TrailerURL:  &trailer3,
			Casting: []models.Actor{
				{ID: 5, FistName: "Jean", LastName: "Dujardin", ImageURL: &img1},
				{ID: 6, FistName: "Marion", LastName: "Cotillard"},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
