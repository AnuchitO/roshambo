package score

import (
	"reflect"
	"testing"
)

func TestNewShouldBeReturnEmtpyScore(t *testing.T) {
	score := &Score{0, 0, 0}

	s := New()

	if !reflect.DeepEqual(s, score) {
		t.Errorf("expect score is zero but got %#v\n", s)
	}
}
