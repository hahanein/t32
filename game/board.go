package game

import "errors"

var (
	ErrIllegalDimensions = errors.New("board: illegal dimensions")
	ErrIllegalMove       = errors.New("board: illegal move")
)

// A Board is a matrix with equal length dimensions. It records the positions
// of all Pieces.
type Board [][]Piece

// validate checks if the Board adheres to the game's specifications and
// returns an error if it is corrupted.
func (b Board) validate() error {
	xSize := len(b)

	for x := 0; x < xSize; x++ {
		ySize := len(b[x])

		if ySize != xSize {
			return ErrIllegalDimensions
		}
	}

	isIllegalSize := xSize < MinSize || xSize > MaxSize

	if isIllegalSize {
		return ErrIllegalDimensions
	}

	return nil
}

// apply applies a Move. If the Move is illegal it returns an error. Otherwise
// it mutates the Board to make it reflect the new state.
func (b Board) apply(ms ...Move) (Board, error) {
	for _, m := range ms {
		ok := b.doesSquareExist(m.X, m.Y)
		if !ok {
			return b, ErrIllegalMove
		}

		ok = b.isSquareEmpty(m.X, m.Y)
		if !ok {
			return b, ErrIllegalMove
		}

		b[m.X][m.Y] = m.Piece
	}

	return b, nil
}

// isSquareEmpty checks if any Piece is occupying a given square and returns
// false if this is the case. Otherwise it returns true.
func (b Board) isSquareEmpty(x, y int) bool {
	return (b)[x][y] == NoPiece
}

// doesSquareExist checks if a square exists at a given position on the Board
// and returns true if this is the case. Otherwise it returns false.
func (b Board) doesSquareExist(x, y int) bool {
	doesXExist := x >= 0 && x < len(b)

	if !doesXExist {
		return false
	}

	doesYExist := y >= 0 && y < len(b[x])

	if !doesYExist {
		return false
	}

	return true
}
