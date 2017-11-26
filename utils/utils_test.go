package utils

import (
	"testing"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func TestUserKeyMapping(t *testing.T) {
	if mapUserKey("w") != decide.Scissors {
		t.Error("expect Scissors but got", mapUserKey("w"))
	}
	if mapUserKey("b") != decide.Rock {
		t.Error("expect Rock but got", mapUserKey("b"))
	}
	if mapUserKey("p") != decide.Paper {
		t.Error("expect Paper but got", mapUserKey("p"))
	}
}
