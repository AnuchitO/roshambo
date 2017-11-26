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

func assertComputerPosition(title string, p int, t *testing.T) {
	c := strings.Index(title, "computer")
	if c != p {
		t.Error("computer should be at position ", p, " but got ", c)
	}
}

func TestTitleShouldBeFixLenghtWhenWIN(t *testing.T) {
	title := Title(decide.WIN)

	assertMePosition(title, t)
	assertComputerPosition(title, 17, t)
}

func TestTitleShouldBeFixLenghtWhenLOSS(t *testing.T) {
	title := Title(decide.LOSS)

	assertMePosition(title, t)
	assertComputerPosition(title, 18, t)
}

func TestTitleShouldBeFixLenghtWhenTIE(t *testing.T) {
	title := Title(decide.TIE)

	assertMePosition(title, t)
	assertComputerPosition(title, 19, t)
}
