package game

import (
	"errors"
)

var (
	ErrIllegalPiece        = errors.New("pieces: has illegal piece")
	ErrDuplicatePieces     = errors.New("pieces: has duplicate pieces")
	ErrWrongNumberOfPieces = errors.New("pieces: has wrong number of pieces")
)

// Piece is the symbol of a given player. A Piece must be unique to a player,
// it must be a printable character and it must be different from the NoPiece
// character.
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

	ok = ps.hasRequiredNumberOfItems()
	if !ok {
		return ErrWrongNumberOfPieces
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

// hasRequiredNumberOfItems returns true if its length reflects the required
// number of players.
func (ps Pieces) hasRequiredNumberOfItems() bool {
	return len(ps) == RequiredNumberOfPlayers
}
