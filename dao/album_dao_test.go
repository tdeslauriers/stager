package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdeslauriers/stager/model"
)

// must create album record in db
func TestCreateAlbum(t *testing.T) {

	test := model.Album{Album: "2018"}
	createAlbum(test)
}

func TestFindAlbumByName(t *testing.T) {

	a := model.Album{Album: "2021"}
	data, _ := createAlbum(a)
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
