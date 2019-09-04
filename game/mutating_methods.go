package game

import (
	"errors"
)

var (
	ErrPlayerExists  = errors.New("player exists")
	ErrPlayerIllegal = errors.New("player illegal")

	ErrGameStarted    = errors.New("game started")
	ErrGameNotStarted = errors.New("game not started")

	ErrMoveNotYourTurn        = errors.New("it is not your turn")
	ErrMoveSquareNotEmpty     = errors.New("square is not empty")
	ErrMoveSquareDoesNotExist = errors.New("square does not exist")
)

// PushPlayer returns a Game with a given Player added. It returns an error if
// the Player must not be added.
//
// Attention: This method mutates its receiver!
func (g *Game) PushPlayer(p Player) error {
	ps, err := g.players.PushPlayer(p)
	if err != nil {
		return err
	}

	g.players = ps

	return nil
}

// PushMove adds a Move to the History. It returns an error if it is illegal or
// if we are still waiting for other Players to join.
//
// Attention: This method mutates its receiver!
func (g *Game) PushMove(m Move) error {
	if g.Status() != StatusRunning {
		return ErrGameNotStarted
	}

	if g.WhoIsNext() != m.Player {
		return ErrMoveNotYourTurn
	}

	ok := g.history.isSquareEmpty(m.X, m.Y)
	if !ok {
		return ErrMoveSquareNotEmpty
	}

	ok = g.size.doesSquareExist(m.X, m.Y)
	if !ok {
		return ErrMoveSquareDoesNotExist
	}

	g.history = append(g.history, m)

	return nil
}

// PushPlayer adds a Player to the list of Players. It returns an error if the
// list is already complete, the Player already exists in the list or if the
// Player is illegal.
func (ps Players) PushPlayer(p Player) (Players, error) {
	if len(ps) >= RequiredNumberOfPlayers {
		return ps, ErrGameStarted
	}

	for i, _ := range ps {
		if ps[i] == p {
			return ps, ErrPlayerExists
		}
	}

	if !p.IsLegal() {
		return ps, ErrPlayerIllegal
	}

	return append(ps, p), nil
}
