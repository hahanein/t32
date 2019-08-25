package participants

import "t32/game"

type spyReferee struct {
	spySubject

	game.Game

	// Set the errors to be returned by their respective function calls.
	ErrRespPushMove   error
	ErrRespGetGame    error
	ErrRespPushPlayer error
}

func (r *spyReferee) PushMove(m game.Move) error {
	if r.ErrRespPushMove != nil {
		return r.ErrRespPushMove
	}

	r.History = append(r.History, m)

	return nil
}

func (r *spyReferee) GetGame() (game.Game, error) {
	if r.ErrRespGetGame != nil {
		return r.Game, r.ErrRespGetGame
	}

	return r.Game, nil
}

func (r *spyReferee) PushPlayer(p game.Player) error {
	if r.ErrRespPushPlayer != nil {
		return r.ErrRespPushPlayer
	}

	r.Players = append(r.Players, p)

	return nil
}
