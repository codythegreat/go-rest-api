package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Artists struct {
	Artists []Artist `json:"artists"`
}
type Artist struct {
	Name  string `json:"Name"`
	Genre string `json:"Genre"`
	ID    int64  `json:"ref"`
}

func getArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func main() {

	jsonFile, err := os.Open("artistInfo.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var artists Artists

	json.Unmarshal(byteValue, &artists)
	for i := 0; i < len(artists.Artists); i++ {
		fmt.Println("Artists: " + artists.Artists[i].Name)
	}
	r := mux.NewRouter()

	r.HandleFunc("/api/artists", getArtists).Methods("GET")
	r.HandleFunc("/api/artists/{id}", getArtists).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
