package utils

import (
	"testing"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func TestShouldBeGreenWhenWIN(t *testing.T) {
	if Color(decide.WIN) != Green {
		t.Error("should be Green but it's not.")
	}
}
func TestShouldBeRedWhenLOSS(t *testing.T) {
	if Color(decide.LOSS) != Red {
		t.Error("should be Red but it's not.")
	}
}

func TestShouldBeYellowWhenTIE(t *testing.T) {
	if Color(decide.TIE) != Yellow {
		t.Error("should be Yellow but it's not.")
	}
}

func TestShouldBeResetColorWhenUnknow(t *testing.T) {
	if Color("") != ResetColor {
		t.Error("should be Reset but it's not.")
	}
}
