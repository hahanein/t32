package participants

type spyCoordinates struct {
	X, Y int
}

type spyClient struct {
	Coordinates []spyCoordinates
	Err         error
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

func (c *spyClient) Fatal(err error) {
	c.Err = err
}
