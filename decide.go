package main

var (
	Paper    = "Paper"
	Scissors = "Scissors"
	Rock     = "Rock"
	TIE      = "TIE"
	WIN      = "WIN"
	LOSS     = "LOSS"
)

func Decide(user, computer string) string {
	if user == Paper {
		if computer == Scissors {
			return LOSS
		}
		if computer == Rock {
			return WIN
		}
	}

	if user == Scissors {
		if computer == Rock {
			return LOSS
		}
		if computer == Paper {
			return WIN
		}
	}

	if user == Rock {
		if computer == Paper {
			return LOSS
		}

		if computer == Scissors {
			return WIN
		}
	}

	return TIE
}
