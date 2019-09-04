package console

import "t32/game"

type spyTemplates struct {
	Dummy string
}

func (t *spyTemplates) WaitingForOthers() string {
	return t.Dummy
}

func (t *spyTemplates) ItsAnothersTurn(b game.Board, p game.Player) string {
	return t.Dummy
}

func (t *spyTemplates) ItsYourTurn(b game.Board, p game.Player) string {
	return t.Dummy
}

func (t *spyTemplates) Stalemate(b game.Board) string {
	return t.Dummy
}

func (t *spyTemplates) AnotherWon(b game.Board, p game.Player) string {
	return t.Dummy
}

func (t *spyTemplates) YouWon(b game.Board, p game.Player) string {
	return t.Dummy
}

func (t *spyTemplates) Flash(b game.Board, msg string) string {
	return t.Dummy
}
