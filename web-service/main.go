package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(albums); err != nil {
		log.Fatal("Error:", err)
	}
}

func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Path[len("/albums/"):]

	for _, album := range albums {
		if album.ID == id {
			if err := json.NewEncoder(w).Encode(album); err != nil {
				log.Fatal("Error:", err)
			}

			return
		}
	}
	http.Error(w, "Album not found", http.StatusNotFound)
}

func postAlbums(w http.ResponseWriter, r *http.Request) {
	var newAlbum Album

	if err := json.NewDecoder(r.Body).Decode(&newAlbum); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newAlbum.ID = strconv.Itoa(len(albums) + 1)
	albums = append(albums, newAlbum)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newAlbum); err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	http.HandleFunc("/albums", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getAlbums(w)
		} else if r.Method == http.MethodPost {
			postAlbums(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/albums/", getAlbumByID)

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
