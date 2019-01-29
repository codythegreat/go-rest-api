package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Artist struct {
	Name  string `json:"Name"`
	Genre string `json:"Genre"`
	ID    string `json:"id"`
}

var jsonArtists []Artist

func printInstructions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Usage:\n/api/artists/all -- get all artists\n/api/artists/ID=123 -- get a specific artists at ID 123\n/api/artists/NAME=Alice-in-Chains -- get instances of artists (- for spaces in name)\n/api/artists/GENRE=vocalists -- get all artists of genre (- for spaces in genre)\n")
}

func getArtists(w http.ResponseWriter, r *http.Request) {
	// set the content type to json format
	w.Header().Set("Content-Type", "application/json")
	// encode all data to page
	json.NewEncoder(w).Encode(jsonArtists)
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	// set the content type to json format
	w.Header().Set("Content-Type", "application/json")
	// extract the parameters of the search
	params := mux.Vars(r)
	for _, item := range jsonArtists {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// for a given artist, retieve all instances
func getArtistGenres(w http.ResponseWriter, r *http.Request) {
	// set the content type to json format
	w.Header().Set("Content-Type", "application/json")
	// extract the parameters of the search
	params := mux.Vars(r)
	var items = []Artist{}
	for _, item := range jsonArtists {
		if item.Name == strings.Join(strings.Split(params["name"], "-"), " ") {
			items = append(items, item)
		}
	}
	json.NewEncoder(w).Encode(items)
}

// for a given genre, retrieve all artist in that genre
func getGenres(w http.ResponseWriter, r *http.Request) {
	// set the content type to json format
	w.Header().Set("Content-Type", "application/json")
	// extract the parameters of the search
	params := mux.Vars(r)
	var items = []Artist{}
	for _, item := range jsonArtists {
		if item.Genre == strings.Join(strings.Split(params["genre"], "-"), " ") {
			items = append(items, item)
		}
	}
	json.NewEncoder(w).Encode(items)
}

func getRandom(w http.ResponseWriter, r *http.Request) {
	// set the content type to json format
	w.Header().Set("Content-Type", "application/json")
	// extract the parameters of the search
	randSrc := rand.NewSource(time.Now().UnixNano())
	randNew := rand.New(randSrc)
	randomPick := randNew.Intn(53383)
	for _, item := range jsonArtists {
		if item.ID == string(randomPick) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	// read the json data from local file
	byteValue, _ := ioutil.ReadFile("artistInfo.json")
	// place the data in jsonArtists
	err := json.Unmarshal(byteValue, &jsonArtists)
	if err != nil {
		fmt.Println(err)
	}
	// initialize the mux router
	r := mux.NewRouter()
	// designate functions to handle routes
	r.HandleFunc("/", printInstructions).Methods("GET")
	r.HandleFunc("/api/artists/all", getArtists).Methods("GET")
	r.HandleFunc("/api/artists/ID={id}", getArtist).Methods("GET")
	r.HandleFunc("/api/artists/NAME={name}", getArtistGenres).Methods("GET")
	r.HandleFunc("/api/artists/GENRE={genre}", getGenres).Methods("GET")
	r.HandleFunc("/api/artists/rand", getRandom).Methods("GET")
	// listen and respond to requests on port 8000
	log.Fatal(http.ListenAndServe("45.76.248.143:80", r))
}
