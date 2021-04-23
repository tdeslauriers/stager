package dao

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tdeslauriers/stager/model"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreateImage(t *testing.T) {

	// missing album id...
	// missing published boolean
	test := model.Pic{Filename: uuid.New(), Date: time.Now(), AlbumID: ObtainAlbumID("2018"), Published: false}
	id, _ := CreateImage(test)
	t.Logf("Created image id: %d", id)
}

func TestFindImageById(t *testing.T) {

	test := model.Pic{Filename: uuid.New(), Date: time.Now(), AlbumID: ObtainAlbumID("2018"), Published: false}
	id, err := CreateImage(test)
	if err != nil {
		t.Log(err)
	}
	f := FindImageById(id)
	t.Log(f.ID, f.Filename, f.Date)
	assert.Equal(t, id, f.ID)

}
