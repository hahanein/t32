package computer

import (
	"t32/game"
)

type Algorithm func(game.Board, game.Player) (int, int)

type Computer struct {
	Algorithm
}

// New returns a new Computer client.
func New(a Algorithm) *Computer {
	c := new(Computer)

	c.Algorithm = a

	return c
}

// WaitingForOthers is called when there need to be more Players before the
// Game may start.
func (c *Computer) WaitingForOthers() {
	// Do nothing.
}

// ItsAnothersTurn is called when it is another Player's turn.
func (c *Computer) ItsAnothersTurn(b game.Board, p game.Player) {
	// Do nothing.
}

// ItsYourTurn is called when it is your turn. You will be prompted to input
// coordinates.
func (c *Computer) ItsYourTurn(b game.Board, p game.Player) (int, int) {
	x, y := c.Algorithm(b, p)

	return x, y
}

// Stalemate is called when there are no more possible Moves but there's also
// no winner.
func (c *Computer) Stalemate(b game.Board) {
	// Do nothing.
}

// YouWon is called when you won the Game.
func (c *Computer) YouWon(b game.Board, p game.Player) {
	// Do nothing.
}

// AnotherWon is called when another Player won the Game.
func (c *Computer) AnotherWon(b game.Board, p game.Player) {
	// Do nothing.
}

// Flash is called when a message is incoming.
func (c *Computer) Flash(b game.Board, msg string) {
	// Do nothing
}
