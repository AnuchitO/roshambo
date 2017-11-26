package score

type Scorer interface {
	IncWin()
	IncLoss()
	IncTie()
}

type Score struct {
	Win  int
	Loss int
	Tie  int
}

func (u *Score) IncWin() {
	u.Win++
}

func (u *Score) IncLoss() {
	u.Loss++
}

func (u *Score) IncTie() {
	u.Tie++
}
