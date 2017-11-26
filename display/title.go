package display

import "github.com/AnuchitPrasertsang/roshambo/decide"

func Title(result string) string {
	me := "    me"
	com := "computer"
	r := map[string]string{
		decide.WIN:  "     win   ",
		decide.LOSS: "     loss   ",
		decide.TIE:  "     tie     ",
	}

	t := r[result]
	if t == "" {
		return me + "             " + com
	}
	return me + t + com
}
