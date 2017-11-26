package play

import (
	"testing"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func TestPlay(t *testing.T) {
	_user := decide.Paper
	computer := decide.Rock
	result := play(_user, computer)
	if result != decide.WIN {
		t.Error("result should be WIN but got", result)
	}
}
