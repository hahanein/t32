package participants

import "t32/game"

type spyReferee struct {
	game.Game
	Err error
}

func (r *spyReferee) PushMove(m game.Move) error {
	r.History = append(r.History, m)
	return nil
}

func (r *spyReferee) GetGame() (game.Game, error) {
	return r.Game, nil
}

func (r *spyReferee) PushPlayer(p game.Player) error {
	r.Players = append(r.Players, p)
	return nil
}

func (r *spyReferee) Fatal(err error) {
	r.Err = err
}
