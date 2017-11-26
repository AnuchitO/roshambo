package computer

import (
	"fmt"
	"testing"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func TestPredictChoice(t *testing.T) {
	s := []string{
		decide.Scissors,
		decide.Paper,
		decide.Scissors,
		decide.Rock,
		decide.Paper,
		decide.Scissors,
		decide.Scissors,
		decide.Paper,
		decide.Rock,
	}
	fmt.Println(s)
}
