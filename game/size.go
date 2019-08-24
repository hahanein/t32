package game

import "errors"

const (
	MaxSize = 10
	MinSize = 3
)

var (
	ErrIllegalSize = errors.New("game: illegal size")
)

type Size int

// doesSquareExist checks if a square exists at a given position on a Board of
// size s and returns true if this is the case. Otherwise it returns false.
func (s Size) doesSquareExist(x, y int) bool {
	return (x >= 0 && x < int(s)) && (y >= 0 && y < int(s))
}

// Validate returns an error if Size s is illegal. Otherwise it returns nil.
func (s Size) Validate() error {
	if s < MinSize || s > MaxSize {
		return ErrIllegalSize
	}

	return nil
}
