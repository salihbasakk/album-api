package main

import (
	"database/sql"
)

func getAlbums(db *sql.DB) ([]Album, error) {
	return getAlbumsFromDB(db)
}

func getAlbumByID(db *sql.DB, id int) (*Album, error) {
	return getAlbumByIDFromDB(db, id)
}

func addAlbum(db *sql.DB, album Album) (int, error) {
	return addAlbumToDB(db, album)
}
