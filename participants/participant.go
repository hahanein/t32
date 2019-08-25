package actors

import (
	"log"
	"t32/actors"
	"t32/game"
)

type Client interface {
	PopCoordinates() (int, int)
	Fatal(error)
}

type Participant struct {
	actors.Referee
	Client
	game.Player
}

// join asks the Referee to join the Game.
func (p *Participant) join() {
	err := p.PushPlayer(p.Player)

	switch err {
	case nil:
		// Do nothing.
	case actors.ErrGameFull:
	case actors.ErrPlayerTaken:
	default:
		log.Fatal(err)
	}
}

// move asks the Referee to accept a Participant's next Move.
func (p *Participant) move() {
	x, y := p.PopCoordinates()

	err := p.PushMove(game.Move{p.Player, x, y})

	switch err {
	case nil:
		// Do nothing.
	case game.ErrIllegalMove:
		p.move()
	default:
		log.Fatal(err)
	}
}
