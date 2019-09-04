// console provides data structures and functions that serve as a bridge
// between a single console user interface and multiple Participants.

package console

import (
	"fmt"
	"io"
	"sync"
	"t32/game"
	"time"
)

type Templates interface {
	WaitingForOthers() string
	ItsAnothersTurn(game.Board, game.Player) string
	ItsYourTurn(game.Board, game.Player) string
	Stalemate(game.Board) string
	AnotherWon(game.Board, game.Player) string
	YouWon(game.Board, game.Player) string
	Flash(game.Board, string) string
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

// New returns a new Console Client.
func New(t Templates, w io.Writer, r io.Reader) *Console {
	c := new(Console)

	c.Templates = t
	c.Writer = w
	c.Reader = r

	return c
}

// WaitingForOthers is called when there need to be more Players before the
// Game may start.
func (c *Console) WaitingForOthers() {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.WaitingForOthers()))
}

// ItsAnothersTurn is called when it is another Player's turn.
func (c *Console) ItsAnothersTurn(b game.Board, p game.Player) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.ItsAnothersTurn(b, p)))
}

// ItsYourTurn is called when it is your turn. You will be prompted to input
// coordinates.
func (c *Console) ItsYourTurn(b game.Board, p game.Player) (int, int) {
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
			c.Write([]byte(c.Templates.Flash(b, err.Error())))
			time.Sleep(1 * time.Second)
			continue
		}

		return x, y
	}
}

// Stalemate is called when there are no more possible Moves but there's also
// no winner.
func (c *Console) Stalemate(b game.Board) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.Stalemate(b)))
}

// YouWon is called when you won the Game.
func (c *Console) YouWon(b game.Board, p game.Player) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.YouWon(b, p)))
}

// AnotherWon is called when another Player won the Game.
func (c *Console) AnotherWon(b game.Board, p game.Player) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.AnotherWon(b, p)))
}

// Flash is called when a message is incoming. It prints a message that may be
// cleared after a second.
func (c *Console) Flash(b game.Board, msg string) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.Flash(b, msg)))
	time.Sleep(1 * time.Second)
}
