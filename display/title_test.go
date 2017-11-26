package display

import (
	"strings"
	"testing"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

var titleLength = 26

func assertMePosition(title string, t *testing.T) {
	m := strings.Index(title, "me")
	if m != 4 {
		t.Error("me should be at position 4 but got ", m)
	}
}

func assertComputerPosition(title string, t *testing.T) {
	c := strings.Index(title, "computer")
	if c != 18 {
		t.Error("computer should be at position 18 but got ", c)
	}
}

func assertTitleLength(title string, t *testing.T) {
	if len(title) != titleLength {
		t.Error("expect tile length ", titleLength, " but got ", len(title))
	}
}

func TestTitleShouldBeFixLenghtWhenWIN(t *testing.T) {
	title := Title(decide.WIN)

	assertTitleLength(title, t)
	assertMePosition(title, t)
	assertComputerPosition(title, t)
}

func TestTitleShouldBeFixLenghtWhenLOSS(t *testing.T) {
	title := Title(decide.LOSS)

	assertTitleLength(title, t)
	assertMePosition(title, t)
	assertComputerPosition(title, t)
}

func TestTitleShouldBeFixLenghtWhenTIE(t *testing.T) {
	title := Title(decide.TIE)

	assertTitleLength(title, t)
	assertMePosition(title, t)
	assertComputerPosition(title, t)
}
