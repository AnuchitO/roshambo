package computer

import (
	"math/rand"
	"time"

	"github.com/AnuchitPrasertsang/roshambo/decide"
)

func Choice() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return choice(rand.Intn(3))
}

func choice(random int) string {
	return decide.Choices[random]
}
