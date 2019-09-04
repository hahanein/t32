package game

import "unicode"

// Player is the symbol of a given participant. A Player must be unique to a
// participant, it must be a printable character and it must be different from
// the NoPlayer character.
type Player rune

type Players []Player

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
