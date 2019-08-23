package entities

const (
	MaxBoardSize    = 10
	MinBoardSize    = 3
	RequiredPlayers = 3
)

type Game struct {
	Pieces
	Board
}

type Pieces []Piece

type Piece rune

var NoPiece Piece
