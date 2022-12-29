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
	ID       string    `json: "id"`
	ISBN     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
}

var movies []Movie

func createMovie(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[:index+1]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func getMovies(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func deleteMovie(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[:index+1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	fmt.Println("Starting Program....")
	movies = append(movies, Movie{ID: "1",
		ISBN:     "999",
		Title:    "The most amazing Panda",
		Director: &Director{FirstName: "Panda", LastName: "Bear"}})
	movies = append(movies, Movie{
		ID:    "2",
		ISBN:  "1000",
		Title: "Panda and Bunny - World Tour",
		Director: &Director{
			FirstName: "Panda",
			LastName:  "Bear",
		},
	})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
