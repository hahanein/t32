package participants

import (
	"t32/actors"
)

type spySubject struct {
	Observers []actors.Observer
}

func (s *spySubject) Attach(o actors.Observer) {
	s.Observers = append(s.Observers, o)
}
