package participant

import "t32/game"

type spyReferee struct {
	spySubject

	NextPlayer game.Player

	game.Players
	game.History
	RespBoard game.Board

	ErrResp error

	RespWinner game.Player
	RespStatus game.Status
}

func (r *spyReferee) PushMove(m game.Move) error {
	if r.ErrResp != nil {
		return r.ErrResp
	}

	r.History = append(r.History, m)

	return nil
}

func (r *spyReferee) Board() game.Board {
	return r.RespBoard
}

func (r *spyReferee) PushPlayer(p game.Player) error {
	if r.ErrResp != nil {
		return r.ErrResp
	}

	r.Players = append(r.Players, p)

	return nil
}

func (r *spyReferee) WhoIsNext() game.Player {
	return r.NextPlayer
}

func (r *spyReferee) Winner() game.Player {
	return r.RespWinner
}

func (r *spyReferee) Status() game.Status {
	return r.RespStatus
}
