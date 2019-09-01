// game contains data structures and functions which encapsulate the complete
// set of rules of "Tic Tac Toe 2.0".
package game

import (
	"errors"
)

const (
	RequiredNumberOfPlayers = 3
)

var (
	ErrSizeIllegal = errors.New("size illegal")
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
	size    Size
	players Players
	history History
}

// New returns a new Game.
func New(s Size) (*Game, error) {
	g := &Game{size: s}

	ok := s.IsLegal()
	if !ok {
		return g, ErrSizeIllegal
	}

	return g, nil
}

// Board derives a Board from the current Game. If Move coordinates in the
// History have duplicates or are out of bounds it returns an error.
func (g *Game) Board() Board {
	b := make(Board, g.size)
	for x, _ := range b {
		b[x] = make([]Player, g.size)
	}

	for _, m := range g.history {
		b[m.X][m.Y] = m.Player
	}

	return b
}

// WhoIsNext derives the Player currently waiting in Line.
func (g *Game) WhoIsNext() (Player, error) {
	if len(g.players) < RequiredNumberOfPlayers {
		return NoPlayer, ErrGameNotStarted
	}

	if len(g.history) == 0 {
		return g.players[0], nil
	}

	i := len(g.history) % len(g.players)

	return g.players[i], nil
}
