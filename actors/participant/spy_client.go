package participant

import "t32/game"

type spyCoordinates struct {
	X, Y int
}

type spyClient struct {
	Coordinates []spyCoordinates
	game.Board
}

func (c *spyClient) PopCoordinates() (int, int) {
	var co spyCoordinates

	if len(c.Coordinates) == 0 {
		return co.X, co.Y
	}

	co, c.Coordinates = c.Coordinates[len(c.Coordinates)-1],
		c.Coordinates[:len(c.Coordinates)-1]

	return co.X, co.Y
}

func (c *spyClient) SetBoard(b game.Board) {
	c.Board = b
}
