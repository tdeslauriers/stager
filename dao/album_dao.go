package dao

import (
	"log"
)

type Album struct {
	ID    int64
	Album string
}

// ObtainAlbumID used to lookup/create album foreign key for Pic
func ObtainAlbumID(name string) (id int64) {

	db := DBConn()
	defer db.Close()

	if a := findAlbumByName(name); a.ID != 0 {
		id = a.ID
		return
	}

	a := Album{Album: name}
	id, _ = InsertAlbum(a)

	db.Close()
	return
}

func InsertAlbum(album Album) (id int64, errSQL error) {

	db := DBConn()
	defer db.Close()

	query := "INSERT INTO album (album) VALUES (?);"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	r, errSQL := stmt.Exec(album.Album)
	if errSQL != nil {
		log.Fatal(errSQL)
	}

	id, errID := r.LastInsertId()
	if errID != nil {
		log.Fatal(errID)
	}

	log.Printf("Created record id: %d\n", id)
	db.Close()

	return id, errSQL
}

func findAlbumByName(name string) (a Album) {

	db := DBConn()
	defer db.Close()

	query := "SELECT id, album FROM album where album = ?"
	row := db.QueryRow(query, name)
	row.Scan(&a.ID, &a.Album)

	db.Close()
	return a
}
