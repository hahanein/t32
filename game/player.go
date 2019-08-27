package game

// Player is the symbol of a given participant. A Player must be unique to a
// participant, it must be a printable character and it must be different from
// the NoPlayer character.
type Player rune

type Players []Player

var NoPlayer Player
