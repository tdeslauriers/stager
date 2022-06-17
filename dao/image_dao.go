package dao

import (
	"log"
	"time"
)

// image table has more fields (title, desc)
type Photo struct {
	ID        int64
	Filename  string
	Date      time.Time
	Published bool
	Thumbnail []byte
	Photo     []byte
}

type Photos []Photo

func FindImageById(id int64) (p Photo) {

	db := DBConn()
	defer db.Close()

	query := "SELECT id, date, BIN_TO_UUID(filename) FROM image WHERE id = ?;"
	row := db.QueryRow(query, id)

	row.Scan(&p.ID, &p.Date, &p.Filename)

	return p
}

func InsertImage(p Photo) (id int64, errSQL error) {

	db := DBConn()
	defer db.Close()

	query := "INSERT INTO image (filename, date, published, thumbnail, image) VALUES (?, ?, ?, ?, ?);"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	r, errSQL := stmt.Exec(p.Filename, p.Date, p.Published, p.Thumbnail, p.Photo)
	if errSQL != nil {
		log.Fatal(errSQL)
	}

	id, errID := r.LastInsertId()
	if errID != nil {
		log.Fatal(errID)
	}

	log.Printf("Created photo-record id: %d\n", id)
	db.Close()

	return id, errSQL
}
