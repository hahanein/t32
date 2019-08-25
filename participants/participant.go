package actors

import (
	"t32/game"
)

type Referee interface {
	PushMove(game.Move) error
	GetGame() (game.Game, error)
	SetPlayer(game.Player) error
	Fatal(error)
}

type Client interface {
	PopPlayer() game.Player
	PopCoordinates() (int, int)
	Fatal(error)
}

type Participant struct {
	Referee
	Client
	game.Player
}

func (p *Participant) isPlayer() bool {
	return p.Player != game.NoPlayer
}

func (p *Participant) joinGame() {
	symbol := p.PopPlayer()

	if symbol == game.NoPlayer {
		return
	}

	err := p.SetPlayer(symbol)
	if err != nil {
		return
	}

	p.Player = symbol
}
