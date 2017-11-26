package utils

import (
	"testing"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func TestUserKeyMapping(t *testing.T) {
	if MapUserKey("w") != decide.Scissors {
		t.Error("expect Scissors but got", MapUserKey("w"))
	}
	if MapUserKey("b") != decide.Rock {
		t.Error("expect Rock but got", MapUserKey("b"))
	}
	if MapUserKey("p") != decide.Paper {
		t.Error("expect Paper but got", MapUserKey("p"))
	}
}
