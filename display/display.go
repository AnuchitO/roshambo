package display

import (
	"strings"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

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

func GetArt(roshambo string) string {
	arts := map[string]string{
		decide.Paper:    PaperArt,
		decide.Rock:     RockArt,
		decide.Scissors: ScissorsArt,
	}
	return arts[roshambo]
}

func Display(userArt, computerArt string) string {
	return concatSameLineForDisplay(splitArtAscii(userArt), splitArtAscii(computerArt))
}
