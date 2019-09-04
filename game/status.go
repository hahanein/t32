package game

type Status int

const (
	StatusWaitingForPlayers Status = iota
	StatusRunning
	StatusFinish
)

// Status derives the current Game status.
func (g *Game) Status() Status {
	if len(g.players) < RequiredNumberOfPlayers {
		return StatusWaitingForPlayers
	}

	if g.Winner() != NoPlayer {
		return StatusFinish
	}

	if g.stalemate() {
		return StatusFinish
	}

	return StatusRunning
}
