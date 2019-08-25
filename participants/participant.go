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

// join asks the Referee to join the Game.
func (p *Participant) join() {
	err := p.PushPlayer(p.Player)
	if err != nil {
		return
	}
}

// move asks the Referee to accept a Participant's next Move.
func (p *Participant) move() {
	x, y := p.PopCoordinates()

	err := p.PushMove(game.Move{p.Player, x, y})
	if err != nil {
		return
	}
}
