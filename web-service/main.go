package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func initDB() (*sql.DB, error) {
	connStr := "user=youruser dbname=yourdb sslmode=disable password=yourpassword host=localhost port=5432"
	return sql.Open("postgres", connStr)
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Failed to close the database:", err)
		}
	}(db)

	http.HandleFunc("GET /albums", func(w http.ResponseWriter, r *http.Request) {
		getAlbumsHandler(db, w, r)
	})
	http.HandleFunc("POST /albums", func(w http.ResponseWriter, r *http.Request) {
		postAlbumsHandler(db, w, r)
	})
	http.HandleFunc("GET /albums/{id}", func(w http.ResponseWriter, r *http.Request) {
		getAlbumByIDHandler(db, w, r)
	})

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
