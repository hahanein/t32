package entities

import (
	"errors"
)

var (
	ErrIllegalPieceInQueue = errors.New("queue: illegal piece")
	ErrIllegalPiece        = errors.New("pieces: illegal piece")
	ErrDuplicatePieces     = errors.New("pieces: duplicate pieces")
)

type Piece rune

type Pieces []Piece

var NoPiece Piece

// MakeCurrentPiece derives the Piece currently waiting in Line from a
// collection of Pieces and a sequence of Moves.
func MakeCurrentPiece(ps Pieces, ms ...Move) (Piece, error) {
	for i := 0; i < len(ms); i++ {
		j := i % len(ps)

		if ps[j] != ms[i].Piece {
			return NoPiece, ErrIllegalPieceInQueue
		}
	}

	i := len(ms) % len(ps)

	return ps[i], nil
}

// validate checks if the list of Pieces adheres to the game's specifications
// and returns an error if it is corrupted.
func (ps Pieces) validate() error {
	for i, _ := range ps {
		if ps[i] == NoPiece {
			return ErrIllegalPiece
		}
	}

	ok := ps.hasUniqItemsOnly()
	if !ok {
		return ErrDuplicatePieces
	}

	return nil
}

// hasUniqItemsOnly returns true if every Piece in the list is different from
// every other Piece in the list. Otherwise it returns false.
func (ps Pieces) hasUniqItemsOnly() bool {
	m := make(map[Piece]struct{})

	for _, p := range ps {
		m[p] = struct{}{}
	}

	if len(ps) != len(m) {
		return false
	}

	return true
}
