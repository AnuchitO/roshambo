package main

import "testing"

// "rock crushes scissors"
func TestRockCrushesScissors(t *testing.T) {
	user := "Rock"
	computer := "Scissors"

	result := Decide(user, computer)

	if "WIN" != result {
		t.Errorf("Rock Crushes Scissors Expected: '%s' but got '%s'", "WIN", result)
	}
}

// "rock covered with paper"
func TestRockLossPaper(t *testing.T) {
	user := "Rock"
	computer := "Paper"

	result := Decide(user, computer)

	if "LOSS" != result {
		t.Errorf("Rock Loss Paper Expected: '%s' but got '%s'", "LOSS", result)
	}
}

// "Rock Tie Rock"
func TestRockTieRock(t *testing.T) {
	user := "Rock"
	computer := "Rock"

	result := Decide(user, computer)

	if "TIE" != result {
		t.Errorf("Rock Tie Rock Expected: '%s' but got '%s'", "TIE", result)
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
