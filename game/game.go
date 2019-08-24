package game

import "errors"

const (
	MaxSize         = 10
	MinSize         = 3
	RequiredPlayers = 3
)

var (
	ErrIllegalPieceInQueue = errors.New("game: illegal piece in queue")
)

// Game is the complete set of data necessary to derive every available piece
// of information about a given Game.
type Game struct {
	Pieces
	Size int
	History
}

// Move represents a Player's intended action. It contains their Piece as well
// as Board coordinates.
type Move struct {
	Piece Piece
	X, Y  int
}

// History is a record of Moves.
type History []Move

// Board derives a Board from the current Game. If Move coordinates in the
// History have duplicates or are out of bounds it returns an error.
func (g Game) Board() (Board, error) {
	b := make(Board, g.Size)
	for x, _ := range b {
		b[x] = make([]Piece, g.Size)
	}

	err := b.validate()
	if err != nil {
		return b, err
	}

	return b.apply(g.History...)
}

// CurrentPiece derives the Piece currently waiting in Line.
func (g Game) CurrentPiece() (Piece, error) {
	for i := 0; i < len(g.History); i++ {
		j := i % len(g.Pieces)

		if g.Pieces[j] != g.History[i].Piece {
			return NoPiece, ErrIllegalPieceInQueue
		}
	}

	i := len(g.History) % len(g.Pieces)

	return g.Pieces[i], nil
}
