package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	user = os.Getenv("GALLERY_TESTDB_GO_USER")
	pass = os.Getenv("GALLERY_TESTDB_GO_PASSWORD")
	dbIP = os.Getenv("GALLERY_TESTDB_GO_IP")
	name = os.Getenv("GALLERY_TESTDB_GO_DBNAME")
)

var url = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, dbIP, name)

// DBConn is db connector function
func DBConn() *sql.DB {
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Printf("Cannot connect to database: %s/%s\n", dbIP, name)
		log.Fatal("Database connection error: ", err)
	} else {
		fmt.Printf("Connected to: %s/%s\n", dbIP, name)
	}
	return db
}
