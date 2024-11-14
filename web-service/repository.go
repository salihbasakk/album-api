package main

import (
	"database/sql"
	"log"
	"sync"
)

var (
	albums      []Album
	albumsMutex sync.Mutex
)

func getAlbumsFromDB(db *sql.DB) ([]Album, error) {
	rows, err := db.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("Error closing rows")
		}
	}(rows)

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, err
		}
		albumsMutex.Lock()
		albums = append(albums, album)
		albumsMutex.Unlock()
	}
	return albums, nil
}

func getAlbumByIDFromDB(db *sql.DB, id int) (*Album, error) {
	var album Album
	err := db.QueryRow("SELECT id, title, artist, price FROM albums WHERE id = $1", id).
		Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &album, err
}

func addAlbumToDB(db *sql.DB, album Album) (int, error) {
	var id int
	err := db.QueryRow(
		"INSERT INTO albums (title, artist, price) VALUES ($1, $2, $3) RETURNING id",
		album.Title, album.Artist, album.Price,
	).Scan(&id)
	return id, err
}
