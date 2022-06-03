package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/barasher/go-exiftool"
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
		}
	}

}
