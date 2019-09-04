package templates

import (
	"fmt"
	"strings"
	"t32/game"
)

const (
	WaitingForOthers = "Waiting for others."
	ItsNobodiesTurn  = "It's nobodies turn?!"
	Stalemate        = "Stalemate..."

	Prompt  = "Enter: "
	Divider = "\n\n"

	// Hungarian "Fmt" prefix for strings which need to be processed with
	// fmt.Sprintf first.
	FmtItsAnothersTurn = "It's %c's turn."
	FmtItsYourTurn     = "It's your turn, %c!"
	FmtYouWon          = "You won, %c!"
	FmtAnotherWon      = "%c won."
)

type Templates struct{}

func (t *Templates) WaitingForOthers() string {
	return WaitingForOthers
}

func (t *Templates) ItsAnothersTurn(b game.Board, p game.Player) string {
	return fmt.Sprintf(FmtItsAnothersTurn, p) +
		Divider +
		Board(b)
}

func (t *Templates) ItsYourTurn(b game.Board, p game.Player) string {
	return fmt.Sprintf(FmtItsYourTurn, p) +
		Divider +
		Board(b) +
		Divider +
		Prompt
}

func (t *Templates) Stalemate(b game.Board) string {
	return Stalemate +
		Divider +
		Board(b)
}

func (t *Templates) YouWon(b game.Board, p game.Player) string {
	return fmt.Sprintf(FmtYouWon, p) +
		Divider +
		Board(b)
}

func (t *Templates) AnotherWon(b game.Board, p game.Player) string {
	return fmt.Sprintf(FmtAnotherWon, p) +
		Divider +
		Board(b)
}

// Board returns the ASCII representation of a Board as a string.
func Board(board game.Board) string {
	var b strings.Builder

	for i := 0; i < len(board); i++ {
		b.WriteString(". ")
	}

	b.WriteString(".\n")

	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] == game.NoPlayer {
				b.WriteString(". ")
			} else {
				b.WriteRune('.')
				b.WriteRune(rune(board[i][j]))
			}
		}

		if i+1 < len(board) {
			b.WriteString(".\n")
		} else {
			b.WriteRune('.')
		}
	}

	return b.String()
}

// Flash returns the current Board accompanied by a flash message.
func (t *Templates) Flash(b game.Board, msg string) string {
	return "Message: " + msg +
		Divider +
		Board(b)
}
