package score

import "testing"

func TestIncrementScoreScoreWin(t *testing.T) {
	u := &Score{}

	u.IncWin()

	if u.Win != 1 {
		t.Error("Score win should be increase")
	}
}

func TestIncrementScoreScoreLoss(t *testing.T) {
	u := &Score{}

	u.IncLoss()

	if u.Loss != 1 {
		t.Error("Score loss should be increase")
	}
}

func TestIncrementScoreScoreTie(t *testing.T) {
	u := &Score{}

	u.IncTie()

	if u.Tie != 1 {
		t.Error("Score loss should be increase")
	}
}
