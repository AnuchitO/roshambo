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
		if computer == Paper {
			return TIE
		}
		return WIN
	}

	if user == Scissors {
		if computer == Rock {
			return LOSS
		}
		if computer == Scissors {
			return TIE
		}
		return WIN
	}

	if computer == Paper {
		return LOSS
	}

	if computer == Rock {
		return TIE
	}

	return WIN
}
