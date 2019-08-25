package actors

type SpyCoordinates struct {
	X, Y int
}

type SpyClient struct {
	Coordinates []SpyCoordinates
	Err         error
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
