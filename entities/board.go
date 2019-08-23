package entities

import "errors"

var (
	ErrUnequalDimensions = errors.New("board: this board has unequal dimensions")
	ErrIllegalSize       = errors.New("board: this size is out of range")
)

// A Board is a matrix with equal length dimensions. It records the positions
// of all Pieces.
type Board [][]Piece

// NewBoard returns a new empty Board of a given size.
func NewBoard(size int) (*Board, error) {
	b := make(Board, size)
	for x, _ := range b {
		b[x] = make([]Piece, size)
	}

	err := b.validate()
	if err != nil {
		return &b, err
	}

	return &b, nil
}

// validate checks if the Board adheres to the game's specifications and
// returns an error if it is corrupted.
func (b *Board) validate() error {
	xSize := len(*b)

	for x := 0; x < xSize; x++ {
		ySize := len((*b)[x])

		if ySize != xSize {
			return ErrUnequalDimensions
		}
	}

	isIllegalSize := xSize < MinBoardSize || xSize > MaxBoardSize

	if isIllegalSize {
		return ErrIllegalSize
	}

	return nil
}
