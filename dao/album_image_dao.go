package dao

import "log"

type AlbumImage struct {
	ID      int64
	AlbumID int64
	PhotoID int64
}

func InsertAlbumImage(a int64, i int64) (id int64, errSQL error) {

	db := DBConn()
	defer db.Close()

	q := "INSERT INTO album_image (album_id, image_id) VALUES (?, ?);"
	s, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	r, errSQL := s.Exec(a, i)
	if errSQL != nil {
		log.Fatal(errSQL)
	}

	id, errID := r.LastInsertId()
	if errID != nil {
		log.Fatal(errID)
	}

	log.Printf("Created xref record: %d, associating album %d with image %d.", id, a, i)
	db.Close()

	return id, errSQL
}
