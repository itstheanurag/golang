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
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}


var movies []Movie

func main() {
	appendMovies()
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting backend server on Port: 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}


func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}


func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(nil)
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie 

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}


func updateMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[:index+1]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)

		}
	}
	json.NewEncoder(w).Encode(movies)
}


func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[:index+1]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}


func appendMovies() {
	movies = append(movies, Movie{
		ID:     "1",
		Isbn:   "438227",
		Title:  "Avengers: Endgame",
		Director: &Director{
			FirstName: "Anthony",
			LastName:  "Russo",
		},
	})
	movies = append(movies, Movie{
		ID:     "2",
		Isbn:   "454555",
		Title:  "Inception",
		Director: &Director{
			FirstName: "Christopher",
			LastName:  "Nolan",
		},
	})
	movies = append(movies, Movie{
		ID:     "3",
		Isbn:   "987654",
		Title:  "Interstellar",
		Director: &Director{
			FirstName: "Christopher",
			LastName:  "Nolan",
		},
	})
	movies = append(movies, Movie{
		ID:     "4",
		Isbn:   "123456",
		Title:  "The Dark Knight",
		Director: &Director{
			FirstName: "Christopher",
			LastName:  "Nolan",
		},
	})
}