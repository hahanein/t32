package actors

import (
	"t32/game"
)

type Referee interface {
	PushMove(game.Move) error
	GetGame() (game.Game, error)
	PushPlayer(game.Player) error
	Fatal(error)
}
