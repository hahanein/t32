package actors

import "t32/game"

type SpyCoordinates struct {
	X, Y int
}

type SpyClient struct {
	Players     []game.Player
	Coordinates []SpyCoordinates
	Err         error
}

func (c *SpyClient) PopPlayer() game.Player {
	if len(c.Players) == 0 {
		return game.NoPlayer
	}

	var symbol game.Player

	symbol, c.Players = c.Players[len(c.Players)-1],
		c.Players[:len(c.Players)-1]

	return symbol
}

func (c *SpyClient) PopCoordinates() (int, int) {
	var co SpyCoordinates

	if len(c.Coordinates) == 0 {
		return co.X, co.Y
	}

	co, c.Coordinates = c.Coordinates[len(c.Coordinates)-1],
		c.Coordinates[:len(c.Coordinates)-1]

	return co.X, co.Y
}

func (c *SpyClient) Fatal(err error) {
	c.Err = err
}
