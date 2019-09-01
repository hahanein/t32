package actors

import (
	"sync"
	"t32/observer"
)

// Subject implements the subject interface of the observer pattern. It is
// responsible for managing a list of Observers and notifying them of state
// changes.
type Subject struct {
	sync.Mutex

	Observers []observer.Observer
}

// Attach appends an Observer to the Subject's registry.
func (s *Subject) Attach(o observer.Observer) {
	s.Lock()
	defer s.Unlock()

	s.Observers = append(s.Observers, o)
}

// Notify asynchronously calls the Update() methods of all of Subject's
// Observers. It's the callees' responsibility to make sure Update() is thread
// safe.
func (s *Subject) Notify() {
	for _, o := range s.Observers {
		go o.Update()
	}
}
