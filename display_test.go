package main

import "testing"

func TestMaxWidthPaperArt(t *testing.T) {
	m := maxWidth(paperArt)

	if m != 53 {
		t.Errorf(" max width should be %d but got %d", 53, m)
	}
}

func TestMaxWidthScissorsArt(t *testing.T) {
	m := maxWidth(scissorsArt)

	if m != 37 {
		t.Errorf(" max width should be %d but got %d", 37, m)
	}
}

func TestMaxWidthRockArt(t *testing.T) {
	m := maxWidth(rockArt)

	if m != 48 {
		t.Errorf(" max width should be %d but got %d", 48, m)
	}
}

func TestFillPadding(t *testing.T) {

}
