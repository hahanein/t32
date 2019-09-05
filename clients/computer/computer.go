// Package computer contains data structures and methods to run an automated
// headless client.

package computer

import (
	"t32/game"
)

// Algorithm is the type of function that may be used to drive the Computer
// Client.
type Algorithm func(game.Board, game.Player) (int, int)

// Computer implements the Client interface. It designed to automatically
// generate inputs and run in headless mode.
type Computer struct {
	Algorithm
}

// New returns a new Computer client.
func New(a Algorithm) *Computer {
	c := new(Computer)

	c.Algorithm = a

	return c
}

// WaitingForOthers does nothing.
func (c *Computer) WaitingForOthers() {
	// Do nothing.
}

// ItsAnothersTurn does nothing.
func (c *Computer) ItsAnothersTurn(b game.Board, p game.Player) {
	// Do nothing.
}

// ItsYourTurn uses the provided algorithm to produce and return new
// coordinates form a given Board and Player.
func (c *Computer) ItsYourTurn(b game.Board, p game.Player) (int, int) {
	x, y := c.Algorithm(b, p)

	return x, y
}

// Stalemate does nothing.
func (c *Computer) Stalemate(b game.Board) {
	// Do nothing.
}

// YouWon does nothing.
func (c *Computer) YouWon(b game.Board, p game.Player) {
	// Do nothing.
}

// AnotherWon does nothing.
func (c *Computer) AnotherWon(b game.Board, p game.Player) {
	// Do nothing.
}

// Flash does nothing.
func (c *Computer) Flash(b game.Board, msg string) {
	// Do nothing
}
