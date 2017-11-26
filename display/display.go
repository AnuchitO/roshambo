package display

import "strings"

func splitArtAscii(art string) []string {
	return strings.Split(art, "\n")
}

func concatSameLineForDisplay(artLeft, artRight []string) string {
	result := []string{}
	for i := 0; i < len(artLeft); i++ {
		result = append(result, artLeft[i]+artRight[i])
	}
	return strings.Join(result, "\n")
}

func Display(user, computer string) string {
	return concatSameLineForDisplay(splitArtAscii(user), splitArtAscii(computer))
}
