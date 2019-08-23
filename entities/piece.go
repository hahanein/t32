package entities

import (
	"errors"
)

var (
	ErrIllegalPiece = errors.New("game: illegal piece")
)

type Pieces []Piece

type Piece rune

var NoPiece Piece

// MakeCurrentPiece derives the Piece currently waiting in Line from a
// collection of Pieces and a sequence of Moves.
func MakeCurrentPiece(ps Pieces, ms ...Move) (Piece, error) {
	for i := 0; i < len(ms); i++ {
		j := i % len(ps)

		if ps[j] != ms[i].Piece {
			return NoPiece, ErrIllegalPiece
		}
	}

	i := len(ms) % len(ps)

	return ps[i], nil
}
