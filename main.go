package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/barasher/go-exiftool"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tdeslauriers/stager/dao"
)

func main() {

	dir := "/home/tombomb/Pictures/stage/"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// rename images, create thumbnails, and create db entries
	// need to test for nil
	imgs := make(dao.Photos, 0, 100)
	for _, f := range files {

		// get imaged created date from exif
		et, err := exiftool.NewExiftool()
		if err != nil {
			log.Fatalf("Error when intializing: %v\n", err)
		}
		defer et.Close()

		var date time.Time
		metadata := et.ExtractMetadata(dir + f.Name())
		for _, datem := range metadata {
			if datem.Err != nil {
				log.Fatalf("Error reading metadata in %v: %v\n", datem.File, datem.Err)
			}
			for k, v := range datem.Fields {
				if k == "DateTimeOriginal" {
					dto, _ := time.Parse("2006:01:02 15:04:05", fmt.Sprint(v))
					date = dto
				}
			}
		}

		year := strconv.Itoa(date.Year())

		// create image + album records

		// insert into db

		// rename files and move to backup dir
		// only if db insert successful

	}

}
