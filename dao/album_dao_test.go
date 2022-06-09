package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// must create album record in db
func TestCreateAlbum(t *testing.T) {

	test := Album{Album: "2018"}
	InsertAlbum(test)
}

func TestFindAlbumByName(t *testing.T) {

	a := Album{Album: "2021"}
	data, _ := InsertAlbum(a)
	test := findAlbumByName("2021")
	assert.Equal(t, data, test.ID)

	nullcase := findAlbumByName("2222")
	t.Log(nullcase)

}

// obtain function
// find by name if exists, return || !exists, create
func TestObtainAlbumId(t *testing.T) {

	d := "2016"
	id := ObtainAlbumID(d)
	t.Log(id)
}
