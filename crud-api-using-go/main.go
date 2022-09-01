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
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// global movie object
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Get movies hit")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
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

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
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

		if item.ID == params["id"] {
			// delete movie
			movies = append(movies[:index], movies[index+1:]...)

			// create new movie with updated data and same id
			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["id"]

			// update it in movies array
			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)

		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello word")

}

func main() {

	r := mux.NewRouter()

	movies = append(movies,
		Movie{ID: "1", ISBN: "125654", Title: "Movie 1", Director: &Director{FirstName: "Gaurav", LastName: "Shinde"}},
		Movie{ID: "2", ISBN: "16511", Title: "Movie 2", Director: &Director{FirstName: "John", LastName: "Doe"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// register router
	http.Handle("/", r)
	fmt.Print("Starting server at 5000")

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
