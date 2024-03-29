package dao

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"os"
	"testing"
	"time"

	"github.com/disintegration/imaging"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// POC
const pic = "/home/atomic/Pictures/stage/e80113dd-0e2b-4ef5-a9cc-19ca63b61d38.jpg"

func TestImageDaoCrud(t *testing.T) {

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
	id, _ := InsertImage(test)
	t.Logf("Inserted image; id: %d", id)

}
