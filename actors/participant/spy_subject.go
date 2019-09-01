package participant

import (
	"t32/observer"
)

type spySubject struct {
	Observers []observer.Observer
}

func (s *spySubject) Attach(o observer.Observer) {
	s.Observers = append(s.Observers, o)
}

func (s *spySubject) Notify() {
	for _, o := range s.Observers {
		go o.Update()
	}
}
