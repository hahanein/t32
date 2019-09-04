package ai

import (
	"math/rand"
	"t32/game"
)

// Random takes a Board and the current Player and returns a random legal Move.
func Random(b game.Board, p game.Player) (int, int) {
	moves := 0
	for x, _ := range b {
		for y, _ := range b[x] {
			if b[x][y] != game.NoPlayer {
				moves++
			}
		}
	}

	numberOfRemainingMoves := (len(b) * len(b)) - moves

	countdown := rand.Intn(numberOfRemainingMoves)

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
