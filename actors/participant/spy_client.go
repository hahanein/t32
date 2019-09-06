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

	ReqWaitingForOthers bool
	ReqItsAnothersTurn  bool
	ReqItsYourTurn      bool
	ReqStalemate        bool
	ReqAnotherWon       bool
	ReqYouWon           bool
	ReqFlash            bool
}

func (c *spyClient) WaitingForOthers() {
	c.ReqWaitingForOthers = true
}

func (c *spyClient) ItsAnothersTurn(b game.Board, p game.Player) {
	c.ReqItsAnothersTurn = true

	c.Board = b
	c.Player = p
}

func (c *spyClient) ItsYourTurn(b game.Board, p game.Player) (int, int) {
	c.ReqItsYourTurn = true

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
	c.ReqStalemate = true
	c.Board = b
}

func (c *spyClient) AnotherWon(b game.Board, p game.Player) {
	c.ReqAnotherWon = true
	c.Board = b
	c.Player = p
}

func (c *spyClient) YouWon(b game.Board, p game.Player) {
	c.ReqYouWon = true
	c.Board = b
	c.Player = p
}

func (c *spyClient) Flash(b game.Board, msg string) {
	c.ReqFlash = true
	c.Board = b
	c.Message = msg
}
