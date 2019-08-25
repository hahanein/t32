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
	SetGame(game.Game)
}

type Participant struct {
	actors.Referee
	Client
	game.Player
}

// Update is the only exported Method of Participant.
func (p *Participant) Update() {
	switch g, err := p.Referee.GetGame(); err {
	case nil:
		p.present(g)

		if g.NextPlayer() == p.Player {
			p.move()
		}
	case game.ErrPlayersMissing:
		p.join()
	default:
		log.Fatal(err)
	}
}

// join asks the Referee to join the Game.
func (p *Participant) join() {
	switch err := p.Referee.PushPlayer(p.Player); err {
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

	switch err := p.Referee.PushMove(game.Move{p.Player, x, y}); err {
	case nil:
		// Do nothing.
	case game.ErrIllegalMove:
		// Do nothing.
	default:
		log.Fatal(err)
	}
}

// present pushes a given Game to the Client for display.
func (p *Participant) present(g game.Game) {
	p.Client.SetGame(g)
}
