// console provides data structures and functions
package console

import (
	"context"
	"fmt"
	"io"
	"sync"
	"t32/game"
)

type Templates interface {
	WaitingForOthers() string
	ItsAnothersTurn(game.Board, game.Player) string
	ItsYourTurn(game.Board, game.Player) string
	Stalemate(game.Board) string
	AnotherWon(game.Board, game.Player) string
	YouWon(game.Board, game.Player) string
}

// Console implements the participant.Client interface. It is responsible
// for presenting the current Game state to Participants and taking user input.
// Multiple Participants may use the same Console.
type Console struct {
	sync.Mutex
	Templates
	io.Writer
	io.Reader
}

func New(t Templates, w io.Writer, r io.Reader) *Console {
	c := new(Console)

	c.Templates = t
	c.Writer = w
	c.Reader = r

	return c
}

// WaitingForOthers is called when there need to be more Players before the
// Game may start.
func (c *Console) WaitingForOthers(ctx context.Context) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.WaitingForOthers()))
}

// ItsAnothersTurn is called when it is another Player's turn.
func (c *Console) ItsAnothersTurn(ctx context.Context, b game.Board, p game.Player) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.ItsAnothersTurn(b, p)))
}

// ItsYourTurn is called when it is your turn. You will be prompted to input
// coordinates.
func (c *Console) ItsYourTurn(ctx context.Context, b game.Board, p game.Player) (int, int) {
	c.Lock()
	defer c.Unlock()
	for {
		c.Write([]byte(c.Templates.ItsYourTurn(b, p)))

		var userInput string

		_, err := fmt.Fscanln(c.Reader, &userInput)
		if err != nil {
			continue
		}

		x, y, err := Parse(userInput)
		if err != nil {
			continue
		}

		return x, y
	}
}

// Stalemate is called when there are no more possible Moves but there's also
// no winner.
func (c *Console) Stalemate(ctx context.Context, b game.Board) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.Stalemate(b)))
}

// YouWon is called when you won the Game.
func (c *Console) YouWon(ctx context.Context, b game.Board, p game.Player) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.YouWon(b, p)))
}

// AnotherWon is called when another Player won the Game.
func (c *Console) AnotherWon(ctx context.Context, b game.Board, p game.Player) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.AnotherWon(b, p)))
}