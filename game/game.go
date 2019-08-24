package game

import (
	"errors"
)

const (
	RequiredNumberOfPlayers = 3
)

var (
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

	err := Validate(g)
	if err != nil {
		return g, err
	}

	return g, nil
}

// Validate determines a game's legality with regards to the Tic Tac Toe 2.0
// specifications and returns an error if they've been broken.
func Validate(g Game) error {
	err := g.Size.Validate()
	if err != nil {
		return err
	}

	err = g.Players.Validate()
	if err != nil {
		return err
	}

	for i, m := range g.History {
		ok := g.History[:i].isSquareEmpty(m.X, m.Y)
		if !ok {
			return ErrIllegalMove
		}

		ok = g.Size.doesSquareExist(m.X, m.Y)
		if !ok {
			return ErrIllegalMove
		}

		ok = g.History[:i].isValidPlayerSequence(g.Players)
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

// NextPlayer derives the Player currently waiting in Line.
func (g Game) NextPlayer() Player {
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
