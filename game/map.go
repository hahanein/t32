package game

import "errors"

var (
	ErrPlayersComplete = errors.New("players complete")
	ErrPlayerExists    = errors.New("player exists")
	ErrPlayerIllegal   = errors.New("player illegal")
)

// PushPlayer returns a Game with a given Player added. It returns an error if
// the Player must not be added.
func (g Game) PushPlayer(p Player) (Game, error) {
	ps, err := g.Players.PushPlayer(p)
	if err != nil {
		return g, err
	}

	return Game{g.Size, ps, g.History}, nil
}

// PushPlayer adds a Player to the list of Players. It returns an error if the
// list is already complete, the Player already exists in the list or if the
// Player is illegal.
func (ps Players) PushPlayer(p Player) (Players, error) {
	if len(ps) >= RequiredNumberOfPlayers {
		return ps, ErrPlayersComplete
	}

	for i, _ := range ps {
		if ps[i] == p {
			return ps, ErrPlayerExists
		}
	}

	if p == NoPlayer {
		return ps, ErrPlayerIllegal
	}

	return append(ps, p), nil
}
