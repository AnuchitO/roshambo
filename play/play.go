package play

import (
	"github.com/AnuchitPrasertsang/roshambo/computer"
	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func Play(user string) (string, string, string) {
	com := computer.Choice()
	result := play(user, com)
	return user, com, result
}

func play(user, com string) (result string) {
	result = decide.Decide(user, com)
	return
}
