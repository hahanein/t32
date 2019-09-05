package game

import "unicode"

// Player is the symbol of a given participant. A Player must be unique to a
// participant, it must be a printable character and it must be different from
// the NoPlayer character.
type Player rune

// Players is represents the list of actual Players in a Game. To this end it
// provides methods to safely mutate its state. It is NOT meant to be used as
// an arbitrary list of Players.
type Players []Player

// NoPlayer is a Player reserved to denote no Player.
var NoPlayer Player

// IsLegal checks a Player character is printable and not one of the reserved
// Players.
func (p Player) IsLegal() bool {
	if p == NoPlayer {
		return false
	}

	if !unicode.IsPrint(rune(p)) {
		return false
	}

	return true
}
