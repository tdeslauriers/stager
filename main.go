package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/tdeslauriers/stager/dao"
	"github.com/tdeslauriers/stager/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dir := "/home/tombomb/Pictures/test/source/"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// rename images, create thumbnails, and create db entries
	// need to test for nil
	imgs := make(model.Pics, 0, 100)
	for _, f := range files {

		// read in data
		p, err := os.Open(dir + f.Name())
		if err != nil {
			log.Panicf("Could not open file: %s.\n%v", f.Name(), err)
		}

		x, err := exif.Decode(p)
		if err != nil {
			log.Panicf("Could not decode image: %s.\n%v", f.Name(), err)
		}

		// data model
		date, _ := x.DateTime()
		year := strconv.Itoa(date.Year())

		pic := model.Pic{Filename: uuid.New(), Date: date, Published: false}
		pic.AlbumID = dao.ObtainAlbumID(year)
		imgs = append(imgs, pic)

		// rename files
		err = os.Rename(dir+f.Name(), dir+pic.Filename.String()+".jpg")
		if err != nil {
			panic(err)
		}

		tmb, _ := x.JpegThumbnail()
		thumb := dir + pic.Filename.String() + "_thumb.jpg"
		makeThumb(tmb, thumb)

		// DAO: only add record to db after rename successful.
		dao.CreateImage(pic)
	}

	//scp images to new web directory

}

func makeThumb(thumbRaw []byte, name string) {

	thumb, _, err := image.Decode(bytes.NewReader(thumbRaw))
	if err != nil {
		panic(err)
	}

	out, _ := os.Create(name)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 30

	err = jpeg.Encode(out, thumb, &opts)

	out.Close()
}
