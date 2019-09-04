package game

type Row []Player

// Winner returns the winning Player if there is one. Otherwise it returns
// NoPlayer.
func (g *Game) Winner() Player {
	b := g.Board()

	for _, r := range b {
		p := r.winner()
		if p != NoPlayer {
			return p
		}
	}

	for _, r := range b.rotate() {
		p := r.winner()
		if p != NoPlayer {
			return p
		}
	}

	p := b.diagonal().winner()
	if p != NoPlayer {
		return p
	}

	p = b.rotate().diagonal().winner()
	if p != NoPlayer {
		return p
	}

	return NoPlayer
}

// stalemate returns true if no more Moves are possible.
func (g *Game) stalemate() bool {
	return int(g.size*g.size) == len(g.history)
}

// winner checks if a Row is occupied by only a single Player (other than
// NoPlayer) and returns that Player otherwise it returns NoPlayer.
func (r Row) winner() Player {
	acc := make(map[Player]struct{})

	for x, _ := range r {
		acc[r[x]] = struct{}{}
	}

	if len(acc) == 1 {
		return r[0]
	} else {
		return NoPlayer
	}
}

// rotate rotates a Board counter clockwise by 90 degrees.
func (b Board) rotate() Board {
	res := make(Board, len(b))

	for i := 0; i < len(b); i++ {
		col := make([]Player, len(b))

		for j := 0; j < len(b); j++ {
			col[j] = b[j][i]
		}

		res[len(b)-i-1] = col
	}

	return res
}

func (b Board) diagonals() (Row, Row) {
	return b.diagonal(), b.rotate().diagonal()
}

// diagonal returns the diagonal spanning from the top left edge of the board
// to the bottom right edge.
func (b Board) diagonal() Row {
	res := make(Row, len(b))

	for i := 0; i < len(b); i++ {
		res[i] = b[i][i]
	}

	return res
}
