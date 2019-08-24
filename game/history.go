package game

// History is a record of Moves.
type History []Move

// isSquareEmpty checks if a given square is empty and returns true if this is
// the case. If any Piece is occupying the square it returns false.
func (h History) isSquareEmpty(x, y int) bool {
	return !h.isSquareOccupied(x, y)
}

// isSquareOccupied checks if a Piece is occupying a given square and returns
// true if this is the case. If it is empty it returns false.
func (h History) isSquareOccupied(x, y int) bool {
	for i, _ := range h {
		if x == h[i].X && y == h[i].Y {
			return true
		}
	}

	return false
}

// isValidPieceSequence returns true if a sequence of Pieces in a History is
// valid given a specific list of Pieces.
func (h History) isValidPieceSequence(ps Pieces) bool {
	for i, _ := range h {
		j := i % len(ps)

		if ps[j] != h[i].Piece {
			return false
		}
	}

	return true
}
