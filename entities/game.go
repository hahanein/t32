package entities

import (
	"errors"
)

const (
	MaxSize         = 10
	MinSize         = 3
	RequiredPlayers = 3
)

var (
	ErrIllegalPiece = errors.New("game: illegal piece")
)

type Game struct {
	Pieces
	History
	Size int
}

type Pieces []Piece

type Piece rune

var NoPiece Piece

type Move struct {
	Piece Piece
	X, Y  int
}

type History []Move

// Apply applies a Move to the Game and returns an error if it is illegal.
// Otherwise it is appended to the History and a copy of the resulting Board is
// returned.
func (g *Game) Apply(m Move) (Board, error) {
	h := append(g.History, m)

	b, err := MakeBoard(g.Size, h...)
	if err != nil {
		return b, err
	}

	if g.NextPiece() != m.Piece {
		return b, ErrIllegalPiece
	}

	g.History = h

	return b, nil
}

// NextPiece returns the next Piece in Line at the current point in the
// History.
func (g *Game) NextPiece() Piece {
	i := len(g.History) % len(g.Pieces)

	return g.Pieces[i]
}
