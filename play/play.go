package play

import (
	"github.com/AnuchitPrasertsang/roshambo/computer"
	"github.com/AnuchitPrasertsang/roshambo/decide"
	"github.com/AnuchitPrasertsang/roshambo/score"
)

type Play struct {
	User     score.Scorer
	Computer score.Scorer
	UserStat []string
}

func New(user, computer score.Scorer) *Play {
	return &Play{user, computer, []string{}}
}

func (p *Play) Play(user string) (string, string, string) {
	p.UserStat = append(p.UserStat, user)
	com := computer.RandomChoice()
	result := play(user, com)
	score.Count(p.User, result, p.Computer)
	return user, com, result
}

func play(user, com string) (result string) {
	result = decide.Decide(user, com)
	return
}
