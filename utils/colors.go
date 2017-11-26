package utils

import "github.com/AnuchitPrasertsang/roshambo/decide"

var (
	Green      = "\033[01;32m"
	Red        = "\033[01;31m"
	Yellow     = "\033[01;33m"
	ResetColor = "\033[00m"
)

func Color(result string) string {
	cm := map[string]string{
		decide.WIN:  Green,
		decide.LOSS: Red,
		decide.TIE:  Yellow,
	}

	c := cm[result]
	if c == "" {
		return ResetColor
	}

	return c
}
