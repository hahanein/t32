// templates contains methods for the plaintext presentation of Game state.

package templates

import (
	"fmt"
	"strings"
	"t32/game"
)

type Templates struct {
	waitingForOthers string
	stalemate        string

	message string
	prompt  string
	divider string

	// Hungarian "fmt" prefix for strings which need to be processed with
	// fmt.Sprintf first.
	fmtItsAnothersTurn string
	fmtItsYourTurn     string
	fmtYouWon          string
	fmtAnotherWon      string
}

func New() *Templates {
	return &Templates{
		waitingForOthers: "Waiting for others.",
		stalemate:        "Stalemate...",

		message: "Message: ",
		prompt:  "Enter: ",
		divider: "\n\n",

		fmtItsAnothersTurn: "It's %c's turn.",
		fmtItsYourTurn:     "It's your turn, %c!",
		fmtYouWon:          "You won, %c!",
		fmtAnotherWon:      "%c won.",
	}
}

func (t *Templates) WaitingForOthers() string {
	return t.waitingForOthers
}

func (t *Templates) ItsAnothersTurn(b game.Board, p game.Player) string {
	return fmt.Sprintf(t.fmtItsAnothersTurn, p) +
		t.divider +
		board(b)
}

func (t *Templates) ItsYourTurn(b game.Board, p game.Player) string {
	return fmt.Sprintf(t.fmtItsYourTurn, p) +
		t.divider +
		board(b) +
		t.divider +
		t.prompt
}

func (t *Templates) Stalemate(b game.Board) string {
	return t.stalemate +
		t.divider +
		board(b)
}

func (t *Templates) YouWon(b game.Board, p game.Player) string {
	return fmt.Sprintf(t.fmtYouWon, p) +
		t.divider +
		board(b)
}

func (t *Templates) AnotherWon(b game.Board, p game.Player) string {
	return fmt.Sprintf(t.fmtAnotherWon, p) +
		t.divider +
		board(b)
}

// board returns the ASCII representation of a board as a string.
func board(board game.Board) string {
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
	return t.message + msg +
		t.divider +
		board(b)
}
