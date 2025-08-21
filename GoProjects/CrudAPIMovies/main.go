package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"math/rand"
	"github.com/gorilla/mux"
)

type Movie struct {
ID string `json:"id"`
Isbn string `json:"isbn"`
Title string `json:"title"`
Actor *Actor `json:"actor"`
}

type Actor struct {
 Fullname string `json:"fullname"`
 Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	for _ , item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	for index , item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index] , movies[index+ 1 : ]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter , r *http.Request){
		w.Header().Set("Content-Type" , "application/json")
		var movie Movie 
		_ = json.NewDecoder(r.Body).Decode(&movie)
		movie.ID = strconv.Itoa(rand.Intn(1000))
		movies = append(movies, movie)
		json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	for index , item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index] , movies[index+ 1 : ]...)
			var movie Movie 
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1" , Isbn: "1000", Title: "DDLJ" , Actor: &Actor{Fullname: "Shahruk" , Lastname: "Khan"} })
	movies = append(movies, Movie{ID: "2" , Isbn: "2000", Title: "Ye Jawani" , Actor: &Actor{Fullname: "Ranveer" , Lastname: "Kapoor"} })

	r.HandleFunc("/movies" , getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies" , createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at PORT 8000")
	err := http.ListenAndServe(":8000" , r)
	if err != nil {
		log.Fatal(err)
	}
}