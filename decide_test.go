package main

import "testing"

// "rock crushes scissors"
func TestRockCrushesScissors(t *testing.T) {
	user := "Rock"
	computer := "Scissors"

	result := Decide(user, computer)

	expect := "WIN"
	if expect != result {
		t.Errorf("%s %s %s but got %s", user, expect, computer, result)
	}
}

// "rock covered with paper"
func TestRockLossPaper(t *testing.T) {
	user := "Rock"
	computer := "Paper"

	result := Decide(user, computer)

	expect := "LOSS"
	if expect != result {
		t.Errorf("%s %s %s but got %s", user, expect, computer, result)
	}
}

// "Rock Tie Rock"
func TestRockTieRock(t *testing.T) {
	user := "Rock"
	computer := "Rock"

	result := Decide(user, computer)

	expect := "TIE"
	if expect != result {
		t.Errorf("%s %s %s but got %s", user, expect, computer, result)
	}
}

// paper covers rock
func TestPaperCoversRock(t *testing.T) {
}

// paper Loss scissors
func TestPaperLossScissors(t *testing.T) {
}

// Paper Tie Paper
func TestPaperTiePaper(t *testing.T) {
}

// "scissors cut[s] paper"
func TestScissorsCutPaper(t *testing.T) {
}

// "scissors loss rock"
func TestScissorsLossRock(t *testing.T) {
}

// "scissors Tie scissors"
func TestScissorsTieScissors(t *testing.T) {
}
