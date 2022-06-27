package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// obtain function
// find by name if exists, return || !exists, create
func TestObtainAlbumId(t *testing.T) {

	albums := [...]string{"2016", "2017", "2020", "2021"}
	for _, v := range albums {

		id := ObtainAlbumID(v)
		t.Log(id)
	}
}

func TestFindAlbumByName(t *testing.T) {

	a := Album{Album: "2021"}
	d := ObtainAlbumID(a.Album)
	test := findAlbumByName("2021")
	assert.Equal(t, d, test.ID)

	nullcase := findAlbumByName("2222")
	t.Log(nullcase)
}
