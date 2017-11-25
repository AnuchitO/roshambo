package display

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplitArtAscii(t *testing.T) {
	a := splitArtAscii(paperArt)

	if !reflect.DeepEqual(a, strings.Split(PaperArt, "\n")) {
		t.Error("split art ascii wrong")
	}
}

func TestMaxWidthPaperArt(t *testing.T) {
	m := maxWidth(splitArtAscii(PaperArt))

	if m != 53 {
		t.Errorf(" max width should be %d but got %d", 53, m)
	}
}

func TestMaxWidthScissorsArt(t *testing.T) {
	m := maxWidth(splitArtAscii(ScissorsArt))

	if m != 37 {
		t.Errorf(" max width should be %d but got %d", 37, m)
	}
}

func TestMaxWidthRockArt(t *testing.T) {
	m := maxWidth(splitArtAscii(RockArt))

	if m != 48 {
		t.Errorf(" max width should be %d but got %d", 48, m)
	}
}
