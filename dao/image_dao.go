package dao

import (
	"fmt"
	"log"

	"github.com/tdeslauriers/stager/model"
)

func FindImageById(id int64) (p model.Pic) {

	db := DBConn()
	defer db.Close()

	query := "SELECT id, date, BIN_TO_UUID(filename) FROM image WHERE id = ?;"
	row := db.QueryRow(query, id)

	row.Scan(&p.ID, &p.Date, &p.Filename)

	return p
}

func CreateImage(pic model.Pic) (id int64, errSQL error) {

	db := DBConn()
	defer db.Close()

	query := "INSERT INTO image (filename, date, published, album_id) VALUES (UUID_TO_BIN(?), ?, ?, ?);"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	r, errSQL := stmt.Exec(pic.Filename, pic.Date, pic.Published, pic.AlbumID)
	if errSQL != nil {
		log.Fatal(errSQL)
	}

	id, errID := r.LastInsertId()
	if errID != nil {
		log.Fatal(errID)
	}

	fmt.Printf("Created record id: %d\n", id)
	db.Close()

	return id, errSQL
}
