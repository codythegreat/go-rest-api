package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//	"os"

	"github.com/gorilla/mux"
)

type Artist struct {
	Name  string `json:"Name"`
	Genre string `json:"Genre"`
	ID    string `json:"id"`
}

var jsonArtists []Artist

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
	r.HandleFunc("/api/artists", getArtists).Methods("GET")
	r.HandleFunc("/api/artists/{id}", getArtist).Methods("GET")
	// listen and respond to requests on port 8000
	log.Fatal(http.ListenAndServe(":8000", r))
}
