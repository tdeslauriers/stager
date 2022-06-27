package dao

import "log"

type AlbumImage struct {
	ID      int64
	AlbumID int64
	PhotoID int64
}

func InsertAlbumImage(ai AlbumImage) (id int64, errSQL error) {

	db := DBConn()
	defer db.Close()

	q := "INSERT INTO album_image (album_id, image_id) VALUES (?, ?);"
	s, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	r, errSQL := s.Exec(ai.AlbumID, ai.PhotoID)
	if errSQL != nil {
		log.Fatal(errSQL)
	}

	id, errID := r.LastInsertId()
	if errID != nil {
		log.Fatal(errID)
	}

	log.Printf("Inserted xref record: %d, associating album %d with image %d.", id, ai.AlbumID, ai.PhotoID)
	db.Close()

	return id, errSQL
}
