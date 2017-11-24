package main

func Decide(user, computer string) string {
	if computer == "Paper" {
		return "LOSS"
	}

	if computer == "Rock" {
		return "TIE"
	}
	return "WIN"
}
