package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"testing"
	"time"

	"github.com/barasher/go-exiftool"
	"github.com/disintegration/imaging"
)

// needs to extract thumbnail if present
func TestReadExif(t *testing.T) {

	// poc; not test code
	et, err := exiftool.NewExiftool()
	if err != nil {
		t.Logf("Error when intializing: %v\n", err)
	}
	defer et.Close()

	metadata := et.ExtractMetadata("/home/atomic/Pictures/stage/e80113dd-0e2b-4ef5-a9cc-19ca63b61d38.jpg")

	for _, datem := range metadata {
		if datem.Err != nil {
			t.Logf("Error concerning %v: %v\n", datem.File, datem.Err)
		}

		// // loop/log all fields
		// for k, v := range datem.Fields {
		// 	t.Logf("[%v] %v\n", k, v)
		// }

		// log relevant fields
		for k, v := range datem.Fields {
			if k == "DateTimeOriginal" {
				t.Logf("%v(%T): %v", k, v, v)
				exifTimeLayout := "2006:01:02 15:04:05" // has to be golang go-live date.
				createDate, _ := time.Parse(exifTimeLayout, fmt.Sprint(v))
				t.Logf("Converted DateTimeOriginal to time.Time: %v", createDate)

			}
			if k == "Orientation" {
				orientation := fmt.Sprint(v)
				t.Logf("Orientation: %s", orientation)
			}
		}
	}

}

const pic = "/home/atomic/Pictures/stage/test1.jpg"

// poc resize
func TestResize(t *testing.T) {

	p, err := os.Open(pic)
	if err != nil {
		t.Log(err)
	}
	defer p.Close()

	img, _, err := image.Decode(p)
	if err != nil {
		t.Log(err)
	}

	thumb := imaging.Resize(img, 0, 200, imaging.Linear)

	n, err := os.Create("/home/atomic/Pictures/stage/thumb.jpg")
	if err != nil {
		t.Log(err)
	}

	opt := jpeg.Options{
		Quality: 90,
	}

	jpeg.Encode(n, thumb, &opt)
}
