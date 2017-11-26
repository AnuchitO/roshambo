package computer

import (
	"math/rand"
	"time"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func RandomChoice() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return randomChoice(rand.Intn(3))
}

func randomChoice(random int) string {
	return decide.Choices[random]
}
