package participant

import (
	"context"
	"log"
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
}

type Referee interface {
	observer.Subject

	PushMove(m game.Move) error
	WhoIsNext() (game.Player, error)
	Finish() (game.Player, bool)
	Board() game.Board
	PushPlayer(p game.Player) error
}

type Participant struct {
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
	ctx, _ := context.WithCancel(context.Background())

	winner, ok := p.Referee.Finish()
	if ok {
		b := p.Referee.Board()

		switch winner {
		case game.NoPlayer:
			p.Client.Stalemate(ctx, b)
		case p.Player:
			p.Client.YouWon(ctx, b, winner)
		default:
			p.Client.AnotherWon(ctx, b, winner)
		}

		p.Done <- struct{}{}
	} else {
		switch player, err := p.Referee.WhoIsNext(); err {
		case nil:
			b := p.Referee.Board()

			if player == p.Player {
				x, y := p.ItsYourTurn(ctx, b, player)

				m := game.Move{p.Player, x, y}

				switch err := p.Referee.PushMove(m); err {
				case nil:
				case game.ErrGameNotStarted:
				case game.ErrMoveNotYourTurn:
					fallthrough
				case game.ErrMoveSquareNotEmpty:
					fallthrough
				case game.ErrMoveSquareDoesNotExist:
					p.Update()
				default:
					log.Fatal(err)
				}
			} else {
				p.Client.ItsAnothersTurn(ctx, b, player)
			}
		case game.ErrGameNotStarted:
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
			log.Fatal(err)
		}
	}
}
