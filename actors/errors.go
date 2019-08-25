package actors

import "errors"

var (
	ErrGameFull    = errors.New("game is full")
	ErrPlayerTaken = errors.New("player is taken")
)
