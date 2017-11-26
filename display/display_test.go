package display

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplitArtAscii(t *testing.T) {
	a := splitArtAscii(PaperArt)

	if !reflect.DeepEqual(a, strings.Split(PaperArt, "\n")) {
		t.Error("split art ascii wrong")
	}
}

func TestHightShouldBeEqual(t *testing.T) {
	pl := len(splitArtAscii(PaperArt))
	rl := len(splitArtAscii(RockArt))
	sl := len(splitArtAscii(ScissorsArt))
	if !(pl == rl && rl == sl) {
		t.Errorf("line Hight should be equal but got paper: %d scissors: %d rock: %d \n", pl, sl, rl)
	}
}

func TestMaxWidthPaperArt(t *testing.T) {
	actual := maxWidth(splitArtAscii(PaperArt))

	expected := 64
	if actual != expected {
		t.Errorf(" max width should be %d but got %d", expected, actual)
	}
}

func TestMaxWidthScissorsArt(t *testing.T) {
	actual := maxWidth(splitArtAscii(ScissorsArt))

	expected := 65
	if actual != expected {
		t.Errorf(" max width should be %d but got %d", expected, actual)
	}
}

func TestMaxWidthRockArt(t *testing.T) {
	actual := maxWidth(splitArtAscii(RockArt))

	expected := 64
	if actual != expected {
		t.Errorf(" max width should be %d but got %d", expected, actual)
	}
}
