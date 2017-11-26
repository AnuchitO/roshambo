package computer

import (
	"testing"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func TestChoice(t *testing.T) {
	if choice(0) != decide.Paper {
		t.Error("choice 0 should be Paper but got", choice(0))
	}
	if choice(1) != decide.Rock {
		t.Error("choice 1 should be Rock but got", choice(1))
	}
	if choice(2) != decide.Scissors {
		t.Error("choice 2 should be Scissors but got", choice(2))
	}
}
