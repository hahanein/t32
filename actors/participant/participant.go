package participant

import (
	"log"
	"sync"
	"t32/game"
	"t32/observer"
)

type Client interface {
	WaitingForOthers()
	ItsAnothersTurn(game.Board, game.Player)
	ItsYourTurn(game.Board, game.Player) (int, int)
	Stalemate(game.Board)
	AnotherWon(game.Board, game.Player)
	YouWon(game.Board, game.Player)
	Flash(game.Board, string)
}

type Referee interface {
	observer.Subject

	PushMove(m game.Move) error
	WhoIsNext() game.Player
	Winner() game.Player
	Status() game.Status
	Board() game.Board
	PushPlayer(p game.Player) error
}

// Participant is a potential Player in the Game but it might just be a
// spectator, too. Since every Observer will receive updates irrespective of
// their success in creating a Player in the Game.
type Participant struct {
	sync.Mutex
	Referee
	Client
	game.Player
	Done chan struct{}
}

// New returns a new Participant.
func New(player game.Player, c Client, r Referee) *Participant {
	p := new(Participant)

	p.Player = player
	p.Client = c
	p.Referee = r
	p.Done = make(chan struct{})

	r.Attach(p)

	p.Update()

	return p
}

// Update is the only exported Method of Participant. It encapsulates the
// complete set of a Participant's reactions to a change of state.
func (p *Participant) Update() {
	p.Lock()

	switch p.Referee.Status() {
	case game.StatusFinish:
		b := p.Referee.Board()

		switch winner := p.Referee.Winner(); winner {
		case game.NoPlayer:
			p.Client.Stalemate(b)
		case p.Player:
			p.Client.YouWon(b, winner)
		default:
			p.Client.AnotherWon(b, winner)
		}

		p.Done <- struct{}{}

	case game.StatusRunning:
		b := p.Referee.Board()

		switch player := p.Referee.WhoIsNext(); player {
		case p.Player:
			x, y := p.ItsYourTurn(b, player)

			m := game.Move{p.Player, x, y}

			switch err := p.Referee.PushMove(m); err {
			case nil:
				// Do nothing.
			case game.ErrMoveNotYourTurn:
				fallthrough
			case game.ErrMoveSquareNotEmpty:
				fallthrough
			case game.ErrMoveSquareDoesNotExist:
				p.Unlock()
				p.Flash(b, err.Error())
				p.Update()
				return
			default:
				log.Fatal(err)
			}
		default:
			p.Client.ItsAnothersTurn(b, player)
		}

	case game.StatusWaitingForPlayers:
		p.Client.WaitingForOthers()

		switch err := p.Referee.PushPlayer(p.Player); err {
		case nil:
			// Do nothing.
		case game.ErrGameStarted:
			// Do nothing.
		case game.ErrPlayerExists:
			// Do nothing.
		case game.ErrPlayerIllegal:
			// Do nothing.
		default:
			log.Fatal(err)
		}

	default:
		log.Fatal("unknown game state")
	}

	p.Unlock()
}
