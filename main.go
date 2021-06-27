package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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

func main() {
	router := mux.NewRouter()
	port := ":1337"

	movies = append(movies, Movie{ID: "1", Isbn: "43826", Title: "Movie 0", Director: &Director{Firstname: "FN 0", Lastname: "LN 0"}})
	movies = append(movies, Movie{ID: "2", Isbn: "43827", Title: "Movie 1", Director: &Director{Firstname: "FN 1", Lastname: "LN 1"}})
	movies = append(movies, Movie{ID: "3", Isbn: "43828", Title: "Movie 2", Director: &Director{Firstname: "FN 2", Lastname: "LN 2"}})
	movies = append(movies, Movie{ID: "4", Isbn: "43829", Title: "Movie 3", Director: &Director{Firstname: "FN 3", Lastname: "LN 3"}})
	movies = append(movies, Movie{ID: "5", Isbn: "43830", Title: "Movie 4", Director: &Director{Firstname: "FN 4", Lastname: "LN 4"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	// set json content type
	writer.Header().Set("Content-Type", "application/json")

	// params
	params := mux.Vars(request)

	// add a new movie (the movie that was sent in the body of postman)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(writer).Encode(movie)
			return
		}
	}
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(movies)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}
