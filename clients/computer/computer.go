package computer

import (
	"math/rand"
	"reflect"
	"t32/game"
	"time"
)

type Computer struct {
	game.Game
}

func (c *Computer) PopCoordinates() (int, int) {
	// Pretend to think hard...
	time.Sleep(time.Duration(rand.Int63n(8)) * time.Second)

	numberOfPossibleMoves := (int(c.Size) * int(c.Size)) - len(c.History)

	countdown := rand.Intn(numberOfPossibleMoves)

	b := c.Board()

	for x, _ := range b {
		for y, _ := range b[x] {
			if b[x][y] == game.NoPlayer && countdown == 0 {
				return x, y
			} else if b[x][y] == game.NoPlayer {
				countdown--
			}
		}
	}

	return -1, -1
}

func (c *Computer) SetGame(g game.Game) {
	if reflect.DeepEqual(c.Game, g) {
		return
	}

	c.Game = g

	// TODO
}
