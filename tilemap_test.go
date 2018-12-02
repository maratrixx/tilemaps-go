package tilemaps

import (
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

/**
type Tilemap struct {
	Type   int
	Startx float64
	Starty float64
	Endx   float64
	Endy   float64
	Level  int
}
*/
func TestBeginMakeTilemap(t *testing.T) {
	tm := Tilemap{
		Type:   2,
		Level:  16,
		Startx: 116.397558,
		Starty: 39.930505,
		Endx:   116.412255,
		Endy:   39.913462,
	}

	img, err := BeginMakeTilemap(tm)

	if err != nil {
		t.Fatal(err.Error())
	}

	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err.Error())
	}

	imgFile, err := os.Create(dir + "/test.png")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = png.Encode(imgFile, img)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log("test succ", dir)

}
