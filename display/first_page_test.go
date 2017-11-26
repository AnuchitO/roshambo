package display

import (
	"reflect"
	"testing"
)

func TestFirstPage(t *testing.T) {
	expected := []string{
		" computer choosed choice",
		" choose your choice",
		" w is scissors",
		" b is rock",
		" p is paper",
		"choose one and enter",
	}
	if !reflect.DeepEqual(FirstPage(), expected) {
		t.Errorf("expect % #v \n but got % #v\n", expected, FirstPage())
	}
}

func TestHint(t *testing.T) {
	expected := " computer should next decide already. choose your next decide. w:scissors, b:rock, p:paper (type `exit` for exit) , chose your choice and press enter:"
	if Hint() != expected {
		t.Errorf("expected %s \n but got %s", expected, Hint())
	}
}
