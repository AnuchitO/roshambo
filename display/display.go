package display

import "strings"

func splitArtAscii(art string) []string {
	return strings.Split(art, "\n")
}

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

func concatSameLineForDisplay(artLeft, artRight []string) string {
	result := []string{}
	for i := 0; i < len(artLeft); i++ {
		result = append(result, artLeft[i]+artRight[i])
	}
	return strings.Join(result, "\n")
}
