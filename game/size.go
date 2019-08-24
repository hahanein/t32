package game

type Size int

// doesSquareExist checks if a square exists at a given position on a Board of
// size s and returns true if this is the case. Otherwise it returns false.
func (s Size) doesSquareExist(x, y int) bool {
	return (x >= 0 && x < int(s)) && (y >= 0 && y < int(s))
}
