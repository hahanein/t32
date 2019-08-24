package game

import (
	"errors"
)

var (
	ErrIllegalPiece    = errors.New("pieces: has illegal piece")
	ErrDuplicatePieces = errors.New("pieces: has duplicate pieces")
	ErrPiecesMissing   = errors.New("pieces: some pieces are missing")
	ErrTooManyPieces   = errors.New("pieces: too many pieces")
)

// Piece is the symbol of a given participant. A Piece must be unique to a
// participant, it must be a printable character and it must be different from
// the NoPiece character.
type Piece rune

type Pieces []Piece

var NoPiece Piece

// Validate checks if the list of Pieces adheres to the game's specifications
// and returns an error if it is corrupted.
func (ps Pieces) Validate() error {
	for i, _ := range ps {
		if ps[i] == NoPiece {
			return ErrIllegalPiece
		}
	}

	ok := ps.hasUniqItemsOnly()
	if !ok {
		return ErrDuplicatePieces
	}

	if len(ps) < RequiredNumberOfPieces {
		return ErrPiecesMissing
	}

	if len(ps) > RequiredNumberOfPieces {
		return ErrTooManyPieces
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

	return len(ps) == len(m)
}
