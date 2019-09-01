package referee

import (
	"log"
	"sync"
	"t32/game"
	"t32/observer"
)

type Referee struct {
	sync.RWMutex
	observer.Subject

	game.Game
}

// NewReferee returns a new Referee.
func NewReferee(g game.Game) *Referee {
	r := new(Referee)

	r.Subject = new(Subject)
	r.Game = g

	return r
}

// PushMove attempts to add a Move to the Game's History. On success it
// notifies all observers of the change of state. Otherwise it returns an error
// to the callee.
func (r *Referee) PushMove(m game.Move) error {
	r.Lock()
	defer r.Unlock()

	err := r.game.PushMove(m)
	if err != nil {
		return err
	}

	r.Notify()

	return nil
}

func (r *Referee) WhoIsNext() (game.Player, error) {
	r.RLock()
	defer r.RUnlock()

	return r.game.WhoIsNext()
}

func (r *Referee) Finish() (game.Player, bool) {
	r.RLock()
	defer r.RUnlock()

	return r.game.Finish()
}

func (r *Referee) Board() game.Board {
	r.RLock()
	defer r.RUnlock()

	return r.game.Board()
}

// PushPlayer adds a Player to the list of players and notifies all observers
// of the change of state. When the push operation fails this methods returns
// an error to the callee.
func (r *Referee) PushPlayer(p game.Player) error {
	r.Lock()
	defer r.Unlock()

	switch err := r.game.PushPlayer(p); err {
	case nil:
		r.Notify()
	case game.ErrGameStarted:
		fallthrough
	case game.ErrPlayerExists:
		fallthrough
	case game.ErrPlayerIllegal:
		return err
	default:
		log.Fatal(err)
	}

	return nil
}
