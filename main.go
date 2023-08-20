package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if params["id"] == item.ID {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if params["id"] == item.ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if params["id"] == item.ID {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "545454",
		Title: "Movie One",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Doe",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "678912",
		Title: "Another Movie",
		Director: &Director{
			Firstname: "Jane",
			Lastname:  "Smith",
		},
	})

	movies = append(movies, Movie{
		ID:    "3",
		Isbn:  "932812",
		Title: "Adventure Film",
		Director: &Director{
			Firstname: "Michael",
			Lastname:  "Johnson",
		},
	})

	movies = append(movies, Movie{
		ID:    "4",
		Isbn:  "123456",
		Title: "Sci-Fi Extravaganza",
		Director: &Director{
			Firstname: "Sarah",
			Lastname:  "Williams",
		},
	})

	movies = append(movies, Movie{
		ID:    "5",
		Isbn:  "789012",
		Title: "Drama Time",
		Director: &Director{
			Firstname: "David",
			Lastname:  "Brown",
		},
	})

	movies = append(movies, Movie{
		ID:    "6",
		Isbn:  "345678",
		Title: "Comedy Nights",
		Director: &Director{
			Firstname: "Emily",
			Lastname:  "Clark",
		},
	})

	movies = append(movies, Movie{
		ID:    "7",
		Isbn:  "901234",
		Title: "Action Packed",
		Director: &Director{
			Firstname: "Robert",
			Lastname:  "Taylor",
		},
	})

	movies = append(movies, Movie{
		ID:    "8",
		Isbn:  "567890",
		Title: "Horror Flick",
		Director: &Director{
			Firstname: "Laura",
			Lastname:  "Lee",
		},
	})

	movies = append(movies, Movie{
		ID:    "9",
		Isbn:  "234567",
		Title: "Romantic Journey",
		Director: &Director{
			Firstname: "Daniel",
			Lastname:  "Martin",
		},
	})

	movies = append(movies, Movie{
		ID:    "10",
		Isbn:  "890123",
		Title: "Mystery Unveiled",
		Director: &Director{
			Firstname: "Olivia",
			Lastname:  "Wright",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}
