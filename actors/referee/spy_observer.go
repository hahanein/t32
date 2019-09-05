package referee

type spyObserver struct {
	ch chan struct{}
}

func (o *spyObserver) Update() {
	o.ch <- struct{}{}
}
