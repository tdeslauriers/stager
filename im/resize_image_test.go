package im

import (
	"image"
	"image/jpeg"
	"os"
	"testing"
)

const pic = "/home/atomic/Pictures/stage/test1.jpg"

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

	thumb := Resize(img)

	n, err := os.Create("/home/atomic/Pictures/stage/thumb.jpg")
	if err != nil {
		t.Log(err)
	}

	opt := jpeg.Options{
		Quality: 90,
	}

	jpeg.Encode(n, thumb, &opt)
}
