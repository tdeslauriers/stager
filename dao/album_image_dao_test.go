package dao

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/barasher/go-exiftool"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAlbumImageCrud(t *testing.T) {

	// create album record
	// get imaged created date, orientation from exif
	et, err := exiftool.NewExiftool()
	if err != nil {
		log.Fatalf("Error when intializing: %v\n", err)
	}
	defer et.Close()

	var date time.Time
	metadata := et.ExtractMetadata(pic)
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

	// create image record
	p, err := os.Open(pic)
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
	_ = jpeg.Encode(tbuf, thumb, &opt)
	dbThumb := tbuf.Bytes()

	// photo to bytes
	buf := new(bytes.Buffer)
	_ = jpeg.Encode(buf, img, &opt)
	dbPhoto := buf.Bytes()

	test := Photo{
		Filename:  uuid.NewString(),
		Date:      time.Now(),
		Published: false,
		Thumbnail: dbThumb,
		Photo:     dbPhoto,
	}
	imageId, _ := InsertImage(test)
	t.Log(imageId)

	var xref AlbumImage
	xref.AlbumID = ObtainAlbumID(strconv.Itoa(date.Year()))
	xref.PhotoID = imageId

	id, _ := InsertAlbumImage(xref)
	assert.NotEqual(t, id, 0)

}
