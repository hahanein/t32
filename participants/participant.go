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

type Client interface {
	PopCoordinates() (int, int)
	Fatal(error)
}

type Participant struct {
	Referee
	Client
	game.Player
}

func (p *Participant) joinGame() {
	err := p.PushPlayer(p.Player)
	if err != nil {
		return
	}
}
