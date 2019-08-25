package participants

// This file contains the Participant data structure. It represents a human or
// a computer that is a potential Player in Game. The following principles must
// be considered when editing this data structure and its methods:
//
// (1) When a Referee method returns an error you may safely assume that the
// game state has not changed AS A RESULT of you calling that method.
// (2) However, you MUST NEVER assume that the last Game state you know of is
// still the most recent state.

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
		// Do nothing.
	case actors.ErrPlayerTaken:
		// Do nothing.
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
		// Do nothing.
	default:
		log.Fatal(err)
	}
}
