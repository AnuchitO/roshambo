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

func Display(userArt, computerArt string) string {
	return concatSameLineForDisplay(splitArtAscii(userArt), splitArtAscii(computerArt))
}
