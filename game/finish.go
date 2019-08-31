package game

type Row []Player

// Finish returns the winning Player if there is one and true if the Game has
// ended. Otherwise it returns false.
func (g *Game) Finish() (Player, bool) {
	b := g.Board()

	for _, r := range b {
		p, ok := r.finish()
		if ok {
			return p, ok
		}
	}

	for _, r := range b.rotate() {
		p, ok := r.finish()
		if ok {
			return p, ok
		}
	}

	p, ok := b.diagonal().finish()
	if ok {
		return p, ok
	}

	p, ok = b.rotate().diagonal().finish()
	if ok {
		return p, ok
	}

	return NoPlayer, g.stalemate()
}

// stalemate returns true if no more Moves are possible.
func (g *Game) stalemate() bool {
	return int(g.size*g.size) == len(g.history)
}

// finish checks if a Row is occupied by only a single Player (other than
// NoPlayer) and returns that Player and true. Otherwise it returns false.
func (r Row) finish() (Player, bool) {
	acc := make(map[Player]struct{})

	for x, _ := range r {
		acc[r[x]] = struct{}{}
	}

	if len(acc) == 1 {
		return r[0], r[0] != NoPlayer
	} else {
		return NoPlayer, false
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
