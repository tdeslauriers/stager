package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/barasher/go-exiftool"
	"github.com/disintegration/imaging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/tdeslauriers/stager/dao"
)

func main() {

	dir := "/home/atomic/Pictures/stage/"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// rename images, create thumbnails, and create db entries
	for _, f := range files {

		// get imaged created date, orientation from exif
		et, err := exiftool.NewExiftool()
		if err != nil {
			log.Fatalf("Error when intializing: %v\n", err)
		}
		defer et.Close()

		var date time.Time
		// var orientation string
		metadata := et.ExtractMetadata(dir + f.Name())
		for _, datem := range metadata {
			if datem.Err != nil {
				log.Fatalf("Error reading metadata in %v: %v\n", datem.File, datem.Err)
			}
			for k, v := range datem.Fields {
				if k == "DateTimeOriginal" {
					date, _ = time.Parse("2006:01:02 15:04:05", fmt.Sprint(v))
				}
				if k == "Orientation" {
					// orientation = fmt.Sprint(v)
				}
			}
		}

		// obtain album record id
		albumId := dao.ObtainAlbumID(strconv.Itoa(date.Year()))

		// create image record
		p, err := os.Open(dir + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer p.Close()

		img, _, err := image.Decode(p)
		if err != nil {
			log.Fatal(err)
		}

		opt := jpeg.Options{
			Quality: 90,
		}

		// create thumb; convert to bytes
		thumb := imaging.Resize(img, 0, 200, imaging.Linear)
		tbuf := new(bytes.Buffer)
		err = jpeg.Encode(tbuf, thumb, &opt)
		dbThumb := tbuf.Bytes()

		// photo to bytes
		buf := new(bytes.Buffer)
		err = jpeg.Encode(buf, img, &opt)
		dbPhoto := buf.Bytes()

		photo := dao.Photo{
			Filename:  uuid.NewString(),
			Date:      date,
			Published: false,
			Thumbnail: dbThumb,
			Photo:     dbPhoto,
		}

		// insert photo record into db; associate with album
		imageId, err := dao.InsertImage(photo)
		if err != nil {
			log.Fatalln(err)
		}

		xref := dao.AlbumImage{AlbumID: albumId, PhotoID: imageId}
		_, err = dao.InsertAlbumImage(xref)
		if err != nil {
			log.Fatal(err)
		}

		// rename files and move to backup dir
		// only if db insert successful

	}

}
