package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func getAlbumsHandler(db *sql.DB, w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	albums, err := getAlbums(db)
	if err != nil {
		http.Error(w, "Failed to retrieve albums", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(albums)
	if err != nil {
		return
	}
}

func getAlbumByIDHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.URL.Path[len("/albums/"):])
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}

	album, err := getAlbumByID(db, id)
	if err != nil {
		http.Error(w, "Failed to retrieve album", http.StatusInternalServerError)
		return
	} else if album == nil {
		http.Error(w, "Album not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(album)
	if err != nil {
		return
	}
}

func postAlbumsHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var newAlbum Album
	if err := json.NewDecoder(r.Body).Decode(&newAlbum); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := addAlbum(db, newAlbum)
	if err != nil {
		http.Error(w, "Failed to insert album", http.StatusInternalServerError)
		return
	}

	newAlbum.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newAlbum)
	if err != nil {
		return
	}
}
