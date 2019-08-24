package game

import (
	"errors"
)

var (
	ErrIllegalPlayer    = errors.New("players: has illegal player")
	ErrDuplicatePlayers = errors.New("players: has duplicate players")
	ErrPlayersMissing   = errors.New("players: some players are missing")
	ErrTooManyPlayers   = errors.New("players: too many players")
)

// Player is the symbol of a given participant. A Player must be unique to a
// participant, it must be a printable character and it must be different from
// the NoPlayer character.
type Player rune

type Players []Player

var NoPlayer Player

// Validate checks if the list of Players adheres to the game's specifications
// and returns an error if it is corrupted.
func (ps Players) Validate() error {
	for i, _ := range ps {
		if ps[i] == NoPlayer {
			return ErrIllegalPlayer
		}
	}

	ok := ps.hasUniqItemsOnly()
	if !ok {
		return ErrDuplicatePlayers
	}

	if len(ps) < RequiredNumberOfPlayers {
		return ErrPlayersMissing
	}

	if len(ps) > RequiredNumberOfPlayers {
		return ErrTooManyPlayers
	}

	return nil
}

// hasUniqItemsOnly returns true if every Player in the list is different from
// every other Player in the list. Otherwise it returns false.
func (ps Players) hasUniqItemsOnly() bool {
	m := make(map[Player]struct{})

	for _, p := range ps {
		m[p] = struct{}{}
	}

	return len(ps) == len(m)
}
