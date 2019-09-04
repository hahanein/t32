package participant

import (
	"context"
	"log"
	"sync"
	"t32/game"
	"t32/observer"
)

type Client interface {
	WaitingForOthers(context.Context)
	ItsAnothersTurn(context.Context, game.Board, game.Player)
	ItsYourTurn(context.Context, game.Board, game.Player) (int, int)
	Stalemate(context.Context, game.Board)
	AnotherWon(context.Context, game.Board, game.Player)
	YouWon(context.Context, game.Board, game.Player)
	Flash(context.Context, game.Board, string)
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

	ctx, _ := context.WithCancel(context.Background())

	switch p.Referee.Status() {
	case game.StatusFinish:
		b := p.Referee.Board()

		switch winner := p.Referee.Winner(); winner {
		case game.NoPlayer:
			p.Client.Stalemate(ctx, b)
		case p.Player:
			p.Client.YouWon(ctx, b, winner)
		default:
			p.Client.AnotherWon(ctx, b, winner)
		}

		p.Done <- struct{}{}

	case game.StatusRunning:
		b := p.Referee.Board()

		switch player := p.Referee.WhoIsNext(); player {
		case p.Player:
			x, y := p.ItsYourTurn(ctx, b, player)

			m := game.Move{p.Player, x, y}

			switch err := p.Referee.PushMove(m); err {
			case nil:
			case game.ErrMoveNotYourTurn:
				fallthrough
			case game.ErrMoveSquareNotEmpty:
				fallthrough
			case game.ErrMoveSquareDoesNotExist:
				p.Unlock()
				p.Flash(ctx, b, err.Error())
				p.Update()
				return
			default:
				log.Fatal(err)
			}
		default:
			p.Client.ItsAnothersTurn(ctx, b, player)
		}

	case game.StatusWaitingForPlayers:
		p.Client.WaitingForOthers(ctx)

		switch err := p.Referee.PushPlayer(p.Player); err {
		case nil:
		case game.ErrGameStarted:
		case game.ErrPlayerExists:
		case game.ErrPlayerIllegal:
		default:
			log.Fatal(err)
		}

	default:
		log.Fatal("unknown game state")
	}

	p.Unlock()
}
