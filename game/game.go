package game

import (
	"errors"
)

const (
	MaxSize = 10
	MinSize = 3

	RequiredNumberOfPlayers = 3
)

var (
	ErrIllegalSize = errors.New("game: illegal size")
	ErrIllegalMove = errors.New("game: illegal move")
)

// Move represents a Player's intended action. It contains their Player as well
// as Board coordinates.
type Move struct {
	Player
	X, Y int
}

// Game is the complete set of data necessary to derive all necessary
// informations about a given Game from.
type Game struct {
	Size
	Players
	History
}

// Make either returns a GUARANTEED LEGAL Tic Tac Toe 2.0 game state or an
// error.
func Make(s Size, ps Players, ms ...Move) (Game, error) {
	g := Game{s, ps, ms}

	err := Validate(g.Size, g.Players, g.History...)
	if err != nil {
		return g, err
	}

	return g, nil
}

func Validate(s Size, ps Players, ms ...Move) error {
	if s < MinSize || s > MaxSize {
		return ErrIllegalSize
	}

	err := ps.Validate()
	if err != nil {
		return err
	}

	var h History = ms
	for i, _ := range h {
		ok := h[:i].isSquareEmpty(h[i].X, h[i].Y)
		if !ok {
			return ErrIllegalMove
		}

		ok = s.doesSquareExist(h[i].X, h[i].Y)
		if !ok {
			return ErrIllegalMove
		}

		ok = h[:i].isValidPlayerSequence(ps)
		if !ok {
			return ErrIllegalMove
		}

	}

	return nil
}

// Board derives a Board from the current Game. If Move coordinates in the
// History have duplicates or are out of bounds it returns an error.
func (g Game) Board() Board {
	b := make(Board, g.Size)
	for x, _ := range b {
		b[x] = make([]Player, g.Size)
	}

	for _, m := range g.History {
		b[m.X][m.Y] = m.Player
	}

	return b
}

// CurrentPlayer derives the Player currently waiting in Line.
func (g Game) CurrentPlayer() Player {
	if len(g.History) == 0 && len(g.Players) > 0 {
		return g.Players[0]
	} else if len(g.History) == 0 {
		// TODO: Maybe we should just Panic since at this point
		// everything must've went wrong.
		return NoPlayer
	}

	i := len(g.History) % len(g.Players)

	return g.Players[i]
}
