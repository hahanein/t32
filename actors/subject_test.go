package actors

import (
	"testing"
)

func TestAttach(t *testing.T) {
	s := new(Subject)
	o := new(SpyObserver)

	before := len(s.Observers)

	s.Attach(o)

	after := len(s.Observers)

	if after != before+1 {
		t.Fatal("subject: failed to attach observer")
	}
}

// TestNotify will raise a fatal error if the done channel blocks due to
// Subject's notify() failing to call its Observers' Update() methods.
func TestNotify(t *testing.T) {
	done := make(chan struct{})

	s := new(Subject)

	n := 3

	for i := 0; i < n; i++ {
		s.Attach(&SpyObserver{done})
	}

	s.Notify()

	for i := 0; i < n; i++ {
		<-done
	}
}
