package main

import (
	"os"
	"testing"

	"github.com/rwcarlsen/goexif/exif"
)

// needs to extract thumbnail if present
func TestMakeThumb(t *testing.T) {

	// poc, not test code
	pic := "/home/tombomb/Pictures/test/thumbtest.jpg"
	p, err := os.Open(pic)
	if err != nil {
		t.Log(err)
	}
	x, err := exif.Decode(p)
	if err != nil {
		t.Log(err)
	}
	thumbByte, err := x.JpegThumbnail()
	if err != nil {
		t.Log(err)
	}
	makeThumb(thumbByte, "/home/tombomb/Pictures/test/test2.jpg")
}
