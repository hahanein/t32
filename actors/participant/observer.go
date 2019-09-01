package participant

import (
	"sync"
)

type Observer struct {
	sync.Mutex

	Participant
}

func (o *Observer) Update() {
	o.Lock()
	defer o.Unlock()
	o.Participant.Update()
}
