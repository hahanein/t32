package game

import (
	"errors"
)

const (
	MaxSize = 10
	MinSize = 3

	RequiredNumberOfPlayers = 3
)

var (
	ErrIllegalSize = errors.New("game: illegal size")
	ErrIllegalMove = errors.New("game: illegal move")
)

// Move represents a Player's intended action. It contains their Piece as well
// as Board coordinates.
type Move struct {
	Piece Piece
	X, Y  int
}

// Game is the complete set of data necessary to derive every available piece
// of information about a given Game.
type Game struct {
	Size int
	Pieces
	History
}

// MakeGame either returns a GUARANTEED LEGAL Tic Tac Toe 2.0 game state or an
// error.
func MakeGame(size int, ps Pieces, ms ...Move) (Game, error) {
	var g Game

	if size < MinSize || size > MaxSize {
		return g, ErrIllegalSize
	}

	g.Size = size

	err := ps.validate()
	if err != nil {
		return g, err
	}

	g.Pieces = ps

	return g.apply(ms...)
}

func (g Game) apply(ms ...Move) (Game, error) {
	h := append(g.History, ms...)

	for i, _ := range h {
		ok := h[:i].isSquareEmpty(h[i].X, h[i].Y)
		if !ok {
			return g, ErrIllegalMove
		}

		ok = g.doesSquareExist(h[i].X, h[i].Y)
		if !ok {
			return g, ErrIllegalMove
		}

		ok = h[:i].isValidPieceSequence(g.Pieces)
		if !ok {
			return g, ErrIllegalMove
		}

	}

	g.History = h

	return g, nil
}

// doesSquareExist checks if a square exists at a given position on the Board
// and returns true if this is the case. Otherwise it returns false.
func (g Game) doesSquareExist(x, y int) bool {
	return (x >= 0 && x < g.Size) && (y >= 0 && y < g.Size)
}

// Board derives a Board from the current Game. If Move coordinates in the
// History have duplicates or are out of bounds it returns an error.
func (g Game) Board() Board {
	b := make(Board, g.Size)
	for x, _ := range b {
		b[x] = make([]Piece, g.Size)
	}

	for _, m := range g.History {
		b[m.X][m.Y] = m.Piece
	}

	return b
}

// CurrentPiece derives the Piece currently waiting in Line.
func (g Game) CurrentPiece() Piece {
	if len(g.History) == 0 && len(g.Pieces) > 0 {
		return g.Pieces[0]
	} else if len(g.History) == 0 {
		// TODO: Maybe we should just Panic since at this point
		// everything must've went wrong.
		return NoPiece
	}

	i := len(g.History) % len(g.Pieces)

	return g.Pieces[i]
}
