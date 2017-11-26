package display

func maxWidth(art []string) int {
	l := 0
	for _, item := range art {
		il := len(item)
		if il > l {
			l = il
		}
	}
	return l
}
