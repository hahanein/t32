package game

const (
	MaxSize = 10
	MinSize = 3
)

type Size int

// doesSquareExist checks if a square exists at a given position on a Board of
// size s and returns true if this is the case. Otherwise it returns false.
func (s Size) doesSquareExist(x, y int) bool {
	return (x >= 0 && x < int(s)) && (y >= 0 && y < int(s))
}

// IsLegal returns false if Size s is illegal. Otherwise it returns true.
func (s Size) IsLegal() bool {
	if s < MinSize || s > MaxSize {
		return false
	}

	return true
}
