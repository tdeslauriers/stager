package im

import (
	"image"
	"os"
	"testing"

	"image/jpeg"
)

const pic = "/home/atomic/Pictures/stage/e80113dd-0e2b-4ef5-a9cc-19ca63b61d38.jpg"

// poc
func TestPoc(t *testing.T) {

	p, err := os.Open(pic)
	if err != nil {
		t.Log(err)
	}
	defer p.Close()

	img, _, err := image.Decode(p)
	if err != nil {
		t.Log(err)
	}

	n, err := os.Create("/home/atomic/Pictures/stage/thumb.jpg")
	if err != nil {
		t.Log(err)
	}

	opt := jpeg.Options{
		Quality: 90,
	}

	jpeg.Encode(n, img, &opt)
}
