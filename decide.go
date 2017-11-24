package main

func Decide(user, computer string) string {
	if user == "Paper" {
		if computer == "Scissors" {
			return "LOSS"
		}
		if computer == "Paper" {
			return "TIE"
		}
		return "WIN"
	}

	if computer == "Paper" {
		return "LOSS"
	}

	if computer == "Rock" {
		return "TIE"
	}

	return "WIN"
}
