// Package console provides data structures and functions that serve as a
// bridge between a single console user interface and multiple Participants.

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

// WaitingForOthers presents a waiting screen to the user.
func (c *Console) WaitingForOthers() {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.WaitingForOthers()))
}

// ItsAnothersTurn presents the Game Board and a message denoting the Player
// that gets to make the next Move.
func (c *Console) ItsAnothersTurn(b game.Board, p game.Player) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.ItsAnothersTurn(b, p)))
}

// ItsYourTurn presents the Game Board and a Message prompting you to insert
// coordinates which will be parsed and returned to the caller.
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

// Stalemate presents the Game Board and a message denoting a stalemate.
func (c *Console) Stalemate(b game.Board) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.Stalemate(b)))
}

// YouWon presents the Game Board and a message about you having won the game.
func (c *Console) YouWon(b game.Board, p game.Player) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.YouWon(b, p)))
}

// AnotherWon presents the Game Board and a message denoting the winning
// Player.
func (c *Console) AnotherWon(b game.Board, p game.Player) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.AnotherWon(b, p)))
}

// Flash presents the Game Board and an arbitrary message to the Player.
func (c *Console) Flash(b game.Board, msg string) {
	c.Lock()
	defer c.Unlock()
	c.Write([]byte(c.Templates.Flash(b, msg)))
	time.Sleep(1 * time.Second)
}
