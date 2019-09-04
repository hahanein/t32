package participant

import "t32/game"

type spyCoordinates struct {
	X, Y int
}

type spyClient struct {
	Coordinates []spyCoordinates
	game.Board
	game.Player
	Message string
}

func (c *spyClient) WaitingForOthers() {
	// Do nothing.
}

func (c *spyClient) ItsAnothersTurn(b game.Board, p game.Player) {
	c.Board = b
	c.Player = p
}

func (c *spyClient) ItsYourTurn(b game.Board, p game.Player) (int, int) {
	c.Board = b
	c.Player = p

	var co spyCoordinates

	if len(c.Coordinates) == 0 {
		return co.X, co.Y
	}

	co, c.Coordinates = c.Coordinates[len(c.Coordinates)-1],
		c.Coordinates[:len(c.Coordinates)-1]

	return co.X, co.Y
}

func (c *spyClient) Stalemate(b game.Board) {
	c.Board = b
}

func (c *spyClient) AnotherWon(b game.Board, p game.Player) {
	c.Board = b
	c.Player = p
}

func (c *spyClient) YouWon(b game.Board, p game.Player) {
	c.Board = b
	c.Player = p
}

func (c *spyClient) Flash(b game.Board, msg string) {
	c.Board = b
	c.Message = msg
}
