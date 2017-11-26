package display

import (
	"reflect"
	"strings"
	"testing"

	"github.com/AnuchitPrasertsang/roshambo/decide"
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

func TestConcatSameLineForDisplay(t *testing.T) {
	r := []string{"rline 1 ", "rline 2 ", "rline 3 "}
	p := []string{"pline 1 ", "pline 2 ", "pline 3 "}

	result := concatSameLineForDisplay(r, p)

	expected := "rline 1 pline 1 \nrline 2 pline 2 \nrline 3 pline 3 "

	if result != expected {
		t.Errorf("expect:\n % #v\n but got: \n % #v\n", expected, result)
	}
}

func TestDisplay(t *testing.T) {
	d := Display(decide.Paper, decide.Rock)

	lenOfConcat := 3742
	if len(d) != lenOfConcat {
		t.Errorf("expect %d, but got '%d'", lenOfConcat, len(d))
	}
}

func TestGetArt(t *testing.T) {
	art := GetArt(decide.Rock)

	if len(art) < 1 {
		t.Error("art is empty")
	}
}
