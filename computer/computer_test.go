package computer

import (
	"testing"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func TestChoice(t *testing.T) {
	if randomChoice(0) != decide.Paper {
		t.Error("choice 0 should be Paper but got", randomChoice(0))
	}
	if randomChoice(1) != decide.Rock {
		t.Error("choice 1 should be Rock but got", randomChoice(1))
	}
	if randomChoice(2) != decide.Scissors {
		t.Error("choice 2 should be Scissors but got", randomChoice(2))
	}
}
