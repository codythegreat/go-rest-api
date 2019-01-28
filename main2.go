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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonArtists)
}

func getArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get parameters for the search
	for _, item := range jsonArtists {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	// open the json file containing the data
	//jsonFile, err := os.Open("artistInfo.json")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer jsonFile.Close()
	// read the
	byteValue, _ := ioutil.ReadFile("artistInfo.json")

	err := json.Unmarshal(byteValue, &jsonArtists)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", len(jsonArtists))
	r := mux.NewRouter()

	r.HandleFunc("/api/artists", getArtists).Methods("GET")
	r.HandleFunc("/api/artists/{id}", getArtist).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
