package decide

var (
	Paper    = "Paper"
	Scissors = "Scissors"
	Rock     = "Rock"
	TIE      = "tie"
	WIN      = "win"
	LOSS     = "loss"
)

func Decide(user, computer string) string {
	paper := map[string]string{
		Paper:    TIE,
		Rock:     WIN,
		Scissors: LOSS,
	}
	rock := map[string]string{
		Paper:    LOSS,
		Rock:     TIE,
		Scissors: WIN,
	}
	scissors := map[string]string{
		Paper:    WIN,
		Rock:     LOSS,
		Scissors: TIE,
	}

	roshambo := map[string]map[string]string{
		Paper:    paper,
		Rock:     rock,
		Scissors: scissors,
	}

	return roshambo[user][computer]
}
