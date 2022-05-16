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
	TITLE    string    `json:"title"`
	DIRECTOR *Director `json:"director"`
}

type Director struct {
	firstName string `json:"firstname"`
	lastName  string `json:"lastname"`
}

var Movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies) // adding the movies slice to the reponse

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, movie := range Movies {
		if params["id"] == movie.ID {
			Movies = append(Movies[:index], Movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(Movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var movieFound bool = false
	for _, movie := range Movies {
		if movie.ID == params["id"] {
			movieFound = true
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
	if movieFound == false {
		json.NewEncoder(w).Encode("movie not found")
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000)) // we just give it a new id

	Movies = append(Movies, movie)

	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, movie := range Movies {
		if movie.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			var newmovie Movie

			_ = json.NewDecoder(r.Body).Decode(&newmovie)
			newmovie.ID = params["id"]
			Movies = append(Movies, newmovie)

			break
		}
	}

	json.NewEncoder(w).Encode(Movies)
}

func main() {
	fmt.Println("Starting the movie server...")

	r := mux.NewRouter()

	movie1 := Movie{ID: "1", ISBN: "12345", TITLE: "Titanic", DIRECTOR: &Director{firstName: "James", lastName: "Stone"}}
	movie2 := Movie{ID: "2", ISBN: "13455", TITLE: "Shawshank Redemption", DIRECTOR: &Director{firstName: "Matt", lastName: "LeBlanc"}}
	movie3 := Movie{ID: "3", ISBN: "45632", TITLE: "Yaariyan", DIRECTOR: &Director{firstName: "Divya", lastName: "Kapoor"}}
	movie4 := Movie{ID: "4", ISBN: "45678", TITLE: "Shershah", DIRECTOR: &Director{firstName: "Manan", lastName: "Sharma"}}

	Movies = append(Movies, movie1)
	Movies = append(Movies, movie2)
	Movies = append(Movies, movie3)
	Movies = append(Movies, movie4)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id", deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
