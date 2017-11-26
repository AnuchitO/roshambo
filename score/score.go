package score

import "github.com/AnuchitPrasertsang/roshambo/decide"

func Count(user Scorer, resultUser string, computer Scorer) {
	if resultUser == decide.WIN {
		user.IncWin()
		computer.IncLoss()
		return
	}

	if resultUser == decide.LOSS {
		user.IncLoss()
		computer.IncWin()
		return
	}

	user.IncTie()
	computer.IncTie()
}
