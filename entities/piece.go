package entities

import (
	"errors"
)

var (
	ErrIllegalPiece    = errors.New("pieces: illegal piece")
	ErrDuplicatePieces = errors.New("pieces: duplicate pieces")
)

// Piece is the symbol of a given player. A Piece is unique.
type Piece rune

type Pieces []Piece

var NoPiece Piece

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
