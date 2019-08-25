package actors

import (
	"t32/game"
)

type Subject interface {
	Attach(Observer)
}

type Observer interface {
	Update()
}

type Referee interface {
	Subject

	PushMove(game.Move) error
	GetGame() (game.Game, error)
	PushPlayer(game.Player) error
}
