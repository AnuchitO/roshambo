package score

import (
	"reflect"
	"testing"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func TestIncreaseWinScore(t *testing.T) {
	userScore := &Score{
		Win:  1,
		Loss: 0,
		Tie:  0,
	}
	comScore := &Score{
		Win:  0,
		Loss: 1,
		Tie:  0,
	}
	user := &Score{}
	computer := &Score{}

	Count(user, decide.WIN, computer)

	if !reflect.DeepEqual(user, userScore) {
		t.Errorf("expect % #v \n but got % #v", userScore, user)
	}
	if !reflect.DeepEqual(computer, comScore) {
		t.Errorf("expect % #v \n but got % #v", comScore, computer)
	}
}

func TestIncreaseLossScore(t *testing.T) {
	userScore := &Score{
		Win:  0,
		Loss: 1,
		Tie:  0,
	}
	comScore := &Score{
		Win:  1,
		Loss: 0,
		Tie:  0,
	}
	user := &Score{}
	computer := &Score{}

	Count(user, decide.LOSS, computer)

	if !reflect.DeepEqual(user, userScore) {
		t.Errorf("expect % #v \n but got % #v", userScore, user)
	}
	if !reflect.DeepEqual(computer, comScore) {
		t.Errorf("expect % #v \n but got % #v", comScore, computer)
	}
}

func TestIncreaseTieScore(t *testing.T) {
	userScore := &Score{
		Win:  0,
		Loss: 0,
		Tie:  1,
	}
	comScore := &Score{
		Win:  0,
		Loss: 0,
		Tie:  1,
	}
	user := &Score{}
	computer := &Score{}

	Count(user, decide.TIE, computer)

	if !reflect.DeepEqual(user, userScore) {
		t.Errorf("expect % #v \n but got % #v", userScore, user)
	}
	if !reflect.DeepEqual(computer, comScore) {
		t.Errorf("expect % #v \n but got % #v", comScore, computer)
	}
}

func TestUnkownResultShouldNotDoAnything(t *testing.T) {

	user := &Score{1, 1, 2}
	computer := &Score{2, 2, 2}

	Count(user, "unkown", computer)

	if !reflect.DeepEqual(user, &Score{1, 1, 2}) {
		t.Errorf("expect user % #v \n but got % #v", &Score{1, 1, 2}, user)
	}
	if !reflect.DeepEqual(computer, &Score{2, 2, 2}) {
		t.Errorf("expect computer % #v \n but got % #v", &Score{2, 2, 2}, computer)
	}
}
